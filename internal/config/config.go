package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/goccy/go-yaml"
)

const (
	DefaultBaseURL = "https://openapi.tossinvest.com"
	EnvClientID    = "TOSSINVEST_CLIENT_ID"
	EnvClientSecret = "TOSSINVEST_CLIENT_SECRET"
	EnvAccountSeq  = "TOSSINVEST_ACCOUNT_SEQ"
	EnvBaseURL     = "TOSSINVEST_BASE_URL"
	EnvAutoConfirm = "TOSSINVEST_AUTO_CONFIRM"
)

type Config struct {
	ClientID     string `yaml:"client_id" toml:"client_id"`
	ClientSecret string `yaml:"client_secret" toml:"client_secret"`
	AccountSeq   *int64 `yaml:"account_seq" toml:"account_seq"`
	BaseURL      string `yaml:"base_url" toml:"base_url"`
}

func Dir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".tossinvest"), nil
}

func Path() (string, error) {
	dir, err := Dir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "config.yaml"), nil
}

func TokenPath() (string, error) {
	dir, err := Dir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "token.json"), nil
}

func Load() (*Config, error) {
	cfg := &Config{BaseURL: DefaultBaseURL}

	path, err := Path()
	if err == nil {
		data, readErr := os.ReadFile(path)
		if readErr == nil {
			if err := yaml.Unmarshal(data, cfg); err != nil {
				return nil, fmt.Errorf("parse config: %w", err)
			}
		} else if !os.IsNotExist(readErr) {
			return nil, readErr
		}
	}

	applyEnv(cfg)
	if cfg.BaseURL == "" {
		cfg.BaseURL = DefaultBaseURL
	}
	return cfg, nil
}

func applyEnv(cfg *Config) {
	if v := os.Getenv(EnvClientID); v != "" {
		cfg.ClientID = v
	}
	if v := os.Getenv(EnvClientSecret); v != "" {
		cfg.ClientSecret = v
	}
	if v := os.Getenv(EnvBaseURL); v != "" {
		cfg.BaseURL = v
	}
	if v := os.Getenv(EnvAccountSeq); v != "" {
		if n, err := strconv.ParseInt(v, 10, 64); err == nil {
			cfg.AccountSeq = &n
		}
	}
}

func (c *Config) MergeOverrides(clientID, clientSecret, baseURL string, accountSeq *int64) {
	if clientID != "" {
		c.ClientID = clientID
	}
	if clientSecret != "" {
		c.ClientSecret = clientSecret
	}
	if baseURL != "" {
		c.BaseURL = baseURL
	}
	if accountSeq != nil {
		c.AccountSeq = accountSeq
	}
}

func (c *Config) ValidateCredentials() error {
	if strings.TrimSpace(c.ClientID) == "" {
		return fmt.Errorf("client_id is required (set via config or %s)", EnvClientID)
	}
	if strings.TrimSpace(c.ClientSecret) == "" {
		return fmt.Errorf("client_secret is required (set via config or %s)", EnvClientSecret)
	}
	return nil
}

func Save(cfg *Config) error {
	dir, err := Dir()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return err
	}
	path, err := Path()
	if err != nil {
		return err
	}
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o600)
}
