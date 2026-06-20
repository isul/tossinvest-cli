package cmd_test

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/isul/tossinvest-cli/internal/cmd"
)

func TestVersionCommand(t *testing.T) {
	app := cmd.NewApp()
	var buf bytes.Buffer
	app.Writer = &buf

	err := app.Run(context.Background(), []string{"tossinvest-cli", "version"})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(buf.String(), cmd.Version) {
		t.Fatalf("output %q missing version", buf.String())
	}
}

func TestConfigPathCommand(t *testing.T) {
	app := cmd.NewApp()
	var buf bytes.Buffer
	app.Writer = &buf

	err := app.Run(context.Background(), []string{"tossinvest-cli", "config", "path"})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(buf.String(), "tossinvest") {
		t.Fatalf("unexpected path: %s", buf.String())
	}
}
