package config_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/isul/tossinvest-cli/internal/config"
)

func TestSaveAndLoad(t *testing.T) {
	dir := t.TempDir()
	t.Setenv("HOME", dir)

	seq := int64(7)
	cfg := &config.Config{
		ClientID:     "c_test",
		ClientSecret: "s_test",
		AccountSeq:   &seq,
		BaseURL:      config.DefaultBaseURL,
	}
	if err := config.Save(cfg); err != nil {
		t.Fatal(err)
	}

	path, err := config.Path()
	if err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(path); err != nil {
		t.Fatal(err)
	}

	loaded, err := config.Load()
	if err != nil {
		t.Fatal(err)
	}
	if loaded.ClientID != "c_test" {
		t.Fatalf("client_id = %q", loaded.ClientID)
	}
	if loaded.AccountSeq == nil || *loaded.AccountSeq != 7 {
		t.Fatalf("account_seq = %v", loaded.AccountSeq)
	}

	wantDir := filepath.Join(dir, ".tossinvest")
	if filepath.Dir(path) != wantDir {
		t.Fatalf("dir = %s want %s", filepath.Dir(path), wantDir)
	}
}
