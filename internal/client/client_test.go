package client_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/isul/tossinvest-cli/internal/client"
	"github.com/isul/tossinvest-cli/internal/config"
)

func TestGetJSON_UnwrapsResult(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/oauth2/token":
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"access_token":"test-token","token_type":"Bearer","expires_in":3600}`))
		case "/api/v1/prices":
			if r.Header.Get("Authorization") != "Bearer test-token" {
				t.Fatalf("missing bearer token")
			}
			_, _ = w.Write([]byte(`{"result":[{"symbol":"005930","price":"70000"}]}`))
		default:
			http.NotFound(w, r)
		}
	}))
	defer srv.Close()

	cfg := &config.Config{
		ClientID:     "c_test",
		ClientSecret: "s_test",
		BaseURL:      srv.URL,
	}
	c := client.New(cfg, false)

	q := url.Values{}
	q.Set("symbols", "005930")
	data, err := c.GetJSON(context.Background(), "/api/v1/prices", q, nil)
	if err != nil {
		t.Fatalf("GetJSON: %v", err)
	}

	var prices []map[string]string
	if err := json.Unmarshal(data, &prices); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if len(prices) != 1 || prices[0]["symbol"] != "005930" {
		t.Fatalf("unexpected data: %s", string(data))
	}
}

func TestGetJSON_AccountHeader(t *testing.T) {
	var gotAccount string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/oauth2/token":
			_, _ = w.Write([]byte(`{"access_token":"tok","token_type":"Bearer","expires_in":3600}`))
		case "/api/v1/holdings":
			gotAccount = r.Header.Get("X-Tossinvest-Account")
			_, _ = w.Write([]byte(`{"result":{"items":[]}}`))
		default:
			http.NotFound(w, r)
		}
	}))
	defer srv.Close()

	cfg := &config.Config{ClientID: "c", ClientSecret: "s", BaseURL: srv.URL}
	c := client.New(cfg, false)
	seq := int64(42)

	_, err := c.GetJSON(context.Background(), "/api/v1/holdings", nil, &seq)
	if err != nil {
		t.Fatal(err)
	}
	if gotAccount != "42" {
		t.Fatalf("account header = %q, want 42", gotAccount)
	}
}

func TestAPIError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/oauth2/token" {
			_, _ = w.Write([]byte(`{"access_token":"tok","expires_in":3600}`))
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error":{"code":"invalid-request","message":"bad request"}}`))
	}))
	defer srv.Close()

	cfg := &config.Config{ClientID: "c", ClientSecret: "s", BaseURL: srv.URL}
	c := client.New(cfg, false)

	_, err := c.GetJSON(context.Background(), "/api/v1/unknown", nil, nil)
	if err == nil {
		t.Fatal("expected error")
	}
	apiErr, ok := err.(*client.APIError)
	if !ok {
		t.Fatalf("expected APIError, got %T", err)
	}
	if apiErr.Code != "invalid-request" {
		t.Fatalf("code = %q", apiErr.Code)
	}
}
