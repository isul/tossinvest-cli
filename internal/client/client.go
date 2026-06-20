package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/isul/tossinvest-cli/internal/config"
)

type Client struct {
	cfg        *config.Config
	httpClient *http.Client
	debug      bool

	tokenMu sync.Mutex
	token   *cachedToken
}

type cachedToken struct {
	AccessToken string    `json:"access_token"`
	TokenType   string    `json:"token_type"`
	ExpiresIn   int       `json:"expires_in"`
	ExpiresAt   time.Time `json:"expires_at"`
}

type APIError struct {
	StatusCode int
	Code       string
	Message    string
	RequestID  string
	Raw        []byte
}

func (e *APIError) Error() string {
	if e.Code != "" {
		return fmt.Sprintf("%s: %s (HTTP %d)", e.Code, e.Message, e.StatusCode)
	}
	return fmt.Sprintf("API error (HTTP %d): %s", e.StatusCode, e.Message)
}

func New(cfg *config.Config, debug bool) *Client {
	return &Client{
		cfg: cfg,
		httpClient: &http.Client{
			Timeout: 60 * time.Second,
		},
		debug: debug,
	}
}

func (c *Client) GetJSON(ctx context.Context, path string, query url.Values, accountSeq *int64) ([]byte, error) {
	if query == nil {
		query = url.Values{}
	}
	return c.do(ctx, http.MethodGet, path, query, nil, accountSeq)
}

func (c *Client) PostJSON(ctx context.Context, path string, body any, accountSeq *int64) ([]byte, error) {
	var payload []byte
	var err error
	if body != nil {
		payload, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}
	return c.do(ctx, http.MethodPost, path, nil, payload, accountSeq)
}

func (c *Client) do(ctx context.Context, method, path string, query url.Values, body []byte, accountSeq *int64) ([]byte, error) {
	var lastErr error
	for attempt := 0; attempt < 4; attempt++ {
		if attempt > 0 {
			wait := time.Duration(attempt) * time.Second
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			case <-time.After(wait):
			}
		}

		token, err := c.accessToken(ctx)
		if err != nil {
			return nil, err
		}

		u, err := url.Parse(c.cfg.BaseURL + path)
		if err != nil {
			return nil, err
		}
		if len(query) > 0 {
			u.RawQuery = query.Encode()
		}

		var bodyReader io.Reader
		if body != nil {
			bodyReader = bytes.NewReader(body)
		}

		req, err := http.NewRequestWithContext(ctx, method, u.String(), bodyReader)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("Accept", "application/json")
		if body != nil {
			req.Header.Set("Content-Type", "application/json")
		}

		seq := accountSeq
		if seq == nil {
			seq = c.cfg.AccountSeq
		}
		if seq != nil {
			req.Header.Set("X-Tossinvest-Account", strconv.FormatInt(*seq, 10))
		}

		if c.debug {
			fmt.Fprintf(os.Stderr, "DEBUG %s %s\n", method, u.String())
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			lastErr = err
			continue
		}

		data, readErr := io.ReadAll(resp.Body)
		resp.Body.Close()
		if readErr != nil {
			return nil, readErr
		}

		if c.debug {
			fmt.Fprintf(os.Stderr, "DEBUG response %d: %s\n", resp.StatusCode, truncate(string(data), 500))
		}

		if resp.StatusCode == http.StatusUnauthorized {
			c.invalidateToken()
			lastErr = parseAPIError(resp.StatusCode, data)
			continue
		}

		if resp.StatusCode == http.StatusTooManyRequests {
			retryAfter := parseRetryAfter(resp)
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			case <-time.After(retryAfter):
			}
			lastErr = parseAPIError(resp.StatusCode, data)
			continue
		}

		if resp.StatusCode >= 400 {
			return nil, parseAPIError(resp.StatusCode, data)
		}

		return unwrapResult(data)
	}
	if lastErr != nil {
		return nil, lastErr
	}
	return nil, fmt.Errorf("request failed after retries")
}

func unwrapResult(data []byte) ([]byte, error) {
	var envelope struct {
		Result json.RawMessage `json:"result"`
		Error  *struct {
			Code    string `json:"code"`
			Message string `json:"message"`
		} `json:"error"`
	}
	if err := json.Unmarshal(data, &envelope); err != nil {
		return data, nil
	}
	if envelope.Error != nil {
		return nil, &APIError{
			Code:    envelope.Error.Code,
			Message: envelope.Error.Message,
			Raw:     data,
		}
	}
	if envelope.Result != nil {
		return envelope.Result, nil
	}
	return data, nil
}

func parseAPIError(status int, data []byte) error {
	var envelope struct {
		Error struct {
			Code      string `json:"code"`
			Message   string `json:"message"`
			RequestID string `json:"requestId"`
		} `json:"error"`
	}
	if err := json.Unmarshal(data, &envelope); err == nil && envelope.Error.Message != "" {
		return &APIError{
			StatusCode: status,
			Code:       envelope.Error.Code,
			Message:    envelope.Error.Message,
			RequestID:  envelope.Error.RequestID,
			Raw:        data,
		}
	}
	return &APIError{StatusCode: status, Message: strings.TrimSpace(string(data)), Raw: data}
}

func parseRetryAfter(resp *http.Response) time.Duration {
	if v := resp.Header.Get("Retry-After"); v != "" {
		if secs, err := strconv.Atoi(v); err == nil {
			return time.Duration(secs) * time.Second
		}
	}
	if v := resp.Header.Get("X-RateLimit-Reset"); v != "" {
		if secs, err := strconv.Atoi(v); err == nil {
			return time.Duration(secs) * time.Second
		}
	}
	return time.Second
}

func (c *Client) accessToken(ctx context.Context) (string, error) {
	c.tokenMu.Lock()
	defer c.tokenMu.Unlock()
	return c.accessTokenLocked(ctx)
}

func (c *Client) accessTokenLocked(ctx context.Context) (string, error) {
	if c.token != nil && time.Now().Before(c.token.ExpiresAt.Add(-30*time.Second)) {
		return c.token.AccessToken, nil
	}

	if tok, err := loadTokenFromDisk(); err == nil && tok != nil {
		if time.Now().Before(tok.ExpiresAt.Add(-30 * time.Second)) {
			c.token = tok
			return tok.AccessToken, nil
		}
	}

	if err := c.cfg.ValidateCredentials(); err != nil {
		return "", err
	}

	form := url.Values{}
	form.Set("grant_type", "client_credentials")
	form.Set("client_id", c.cfg.ClientID)
	form.Set("client_secret", c.cfg.ClientSecret)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.cfg.BaseURL+"/oauth2/token", strings.NewReader(form.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode >= 400 {
		return "", parseAPIError(resp.StatusCode, data)
	}

	var tok cachedToken
	if err := json.Unmarshal(data, &tok); err != nil {
		return "", err
	}
	if tok.ExpiresIn > 0 {
		tok.ExpiresAt = time.Now().Add(time.Duration(tok.ExpiresIn) * time.Second)
	} else {
		tok.ExpiresAt = time.Now().Add(24 * time.Hour)
	}
	c.token = &tok
	_ = saveTokenToDisk(&tok)
	return tok.AccessToken, nil
}

func (c *Client) invalidateToken() {
	c.token = nil
	_ = removeTokenFromDisk()
}

func (c *Client) IssueToken(ctx context.Context) (*cachedToken, error) {
	c.tokenMu.Lock()
	defer c.tokenMu.Unlock()
	c.token = nil
	_ = removeTokenFromDisk()
	if _, err := c.accessTokenLocked(ctx); err != nil {
		return nil, err
	}
	return c.token, nil
}

func loadTokenFromDisk() (*cachedToken, error) {
	path, err := config.TokenPath()
	if err != nil {
		return nil, err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var tok cachedToken
	if err := json.Unmarshal(data, &tok); err != nil {
		return nil, err
	}
	return &tok, nil
}

func saveTokenToDisk(tok *cachedToken) error {
	path, err := config.TokenPath()
	if err != nil {
		return err
	}
	dir, err := config.Dir()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return err
	}
	data, err := json.Marshal(tok)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o600)
}

func removeTokenFromDisk() error {
	path, err := config.TokenPath()
	if err != nil {
		return err
	}
	err = os.Remove(path)
	if os.IsNotExist(err) {
		return nil
	}
	return err
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "..."
}

func (c *Client) ResolveAccountSeq(ctx context.Context, override *int64) (*int64, error) {
	if override != nil {
		return override, nil
	}
	if c.cfg.AccountSeq != nil {
		return c.cfg.AccountSeq, nil
	}

	data, err := c.GetJSON(ctx, "/api/v1/accounts", nil, nil)
	if err != nil {
		return nil, fmt.Errorf("resolve account: %w", err)
	}
	var accounts []struct {
		AccountSeq int64 `json:"accountSeq"`
	}
	if err := json.Unmarshal(data, &accounts); err != nil {
		return nil, err
	}
	if len(accounts) == 0 {
		return nil, fmt.Errorf("no accounts found")
	}
	if len(accounts) == 1 {
		seq := accounts[0].AccountSeq
		return &seq, nil
	}
	return nil, fmt.Errorf("multiple accounts found; set account_seq in config or use --account-seq")
}
