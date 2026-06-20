package cmd

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/isul/tossinvest-cli/internal/config"
	"github.com/urfave/cli/v3"
)

func configCommand(r *Runtime) *cli.Command {
	return &cli.Command{
		Name:  "config",
		Usage: "Manage CLI configuration",
		Commands: []*cli.Command{
			{
				Name:  "set",
				Usage: "Set credentials and default account interactively",
				Action: func(_ context.Context, _ *cli.Command) error {
					reader := bufio.NewReader(os.Stdin)
					cfg, err := config.Load()
					if err != nil {
						return err
					}

					fmt.Print("Client ID: ")
					clientID, _ := reader.ReadString('\n')
					if v := strings.TrimSpace(clientID); v != "" {
						cfg.ClientID = v
					}

					fmt.Print("Client Secret: ")
					clientSecret, _ := reader.ReadString('\n')
					if v := strings.TrimSpace(clientSecret); v != "" {
						cfg.ClientSecret = v
					}

					fmt.Print("Account Seq (optional): ")
					accountLine, _ := reader.ReadString('\n')
					if v := strings.TrimSpace(accountLine); v != "" {
						n, err := strconv.ParseInt(v, 10, 64)
						if err != nil {
							return fmt.Errorf("invalid account_seq: %w", err)
						}
						cfg.AccountSeq = &n
					}

					fmt.Printf("Base URL [%s]: ", config.DefaultBaseURL)
					baseURL, _ := reader.ReadString('\n')
					if v := strings.TrimSpace(baseURL); v != "" {
						cfg.BaseURL = v
					} else if cfg.BaseURL == "" {
						cfg.BaseURL = config.DefaultBaseURL
					}

					if err := config.Save(cfg); err != nil {
						return err
					}
					path, _ := config.Path()
					fmt.Fprintf(os.Stderr, "Configuration saved to %s\n", path)
					return nil
				},
			},
			{
				Name:  "path",
				Usage: "Show config file path",
				Action: func(_ context.Context, c *cli.Command) error {
					path, err := config.Path()
					if err != nil {
						return err
					}
					_, err = fmt.Fprintln(c.Root().Writer, path)
					return err
				},
			},
		},
	}
}

func authCommand(r *Runtime) *cli.Command {
	return &cli.Command{
		Name:  "auth",
		Usage: "Authentication utilities",
		Commands: []*cli.Command{
			{
				Name:  "token",
				Usage: "Issue OAuth access token (debug)",
				Action: func(ctx context.Context, _ *cli.Command) error {
					tok, err := r.Client.IssueToken(ctx)
					if err != nil {
						return err
					}
					return r.emitJSON(map[string]any{
						"access_token": tok.AccessToken,
						"token_type":   tok.TokenType,
						"expires_in":   tok.ExpiresIn,
						"expires_at":   tok.ExpiresAt,
					})
				},
			},
		},
	}
}
