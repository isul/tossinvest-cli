package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/isul/tossinvest-cli/internal/client"
	"github.com/isul/tossinvest-cli/internal/config"
	"github.com/isul/tossinvest-cli/internal/output"
	"github.com/urfave/cli/v3"
)

const Version = "0.1.0"

type Runtime struct {
	Cfg           *config.Config
	Client        *client.Client
	Output        output.Options
	Debug         bool
	ClientID      string
	ClientSecret  string
	BaseURL       string
	AccountSeq    int64
	AccountSeqSet bool
	Yes           bool
}

func (r *Runtime) Init() error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}
	var accountSeq *int64
	if r.AccountSeqSet {
		accountSeq = &r.AccountSeq
	}
	cfg.MergeOverrides(r.ClientID, r.ClientSecret, r.BaseURL, accountSeq)
	r.Cfg = cfg
	r.Client = client.New(cfg, r.Debug)
	return nil
}

func (r *Runtime) accountSeq(ctx context.Context) (*int64, error) {
	var override *int64
	if r.AccountSeqSet {
		override = &r.AccountSeq
	}
	return r.Client.ResolveAccountSeq(ctx, override)
}

func (r *Runtime) emit(data []byte) error {
	return output.Write(data, r.Output)
}

func (r *Runtime) emitErr(err error) error {
	return output.WriteError(err, r.Output)
}

func (r *Runtime) emitJSON(v any) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return r.emit(data)
}

func globalFlags(r *Runtime) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "client-id",
			Usage:       "OAuth client ID",
			Sources:     cli.EnvVars(config.EnvClientID),
			Destination: &r.ClientID,
		},
		&cli.StringFlag{
			Name:        "client-secret",
			Usage:       "OAuth client secret",
			Sources:     cli.EnvVars(config.EnvClientSecret),
			Destination: &r.ClientSecret,
		},
		&cli.Int64Flag{
			Name:    "account-seq",
			Usage:   "Account sequence for X-Tossinvest-Account header",
			Sources: cli.EnvVars(config.EnvAccountSeq),
			Action: func(_ context.Context, _ *cli.Command, v int64) error {
				r.AccountSeq = v
				r.AccountSeqSet = true
				return nil
			},
		},
		&cli.StringFlag{
			Name:        "base-url",
			Usage:       "API base URL",
			Value:       config.DefaultBaseURL,
			Sources:     cli.EnvVars(config.EnvBaseURL),
			Destination: &r.BaseURL,
		},
		&cli.StringFlag{
			Name:  "format",
			Usage: "Output format: auto, json, yaml, pretty, raw",
			Value: "auto",
			Action: func(_ context.Context, _ *cli.Command, v string) error {
				f, err := output.ParseFormat(v)
				if err != nil {
					return err
				}
				r.Output.Format = f
				return nil
			},
		},
		&cli.StringFlag{
			Name:  "format-error",
			Usage: "Error output format",
			Action: func(_ context.Context, _ *cli.Command, v string) error {
				f, err := output.ParseFormat(v)
				if err != nil {
					return err
				}
				r.Output.FormatError = f
				return nil
			},
		},
		&cli.StringFlag{
			Name:        "transform",
			Usage:       "GJSON transform path",
			Destination: &r.Output.Transform,
		},
		&cli.BoolFlag{
			Name:        "debug",
			Usage:       "Enable debug logging",
			Destination: &r.Debug,
		},
		&cli.BoolFlag{
			Name:        "yes",
			Usage:       "Skip CONFIRM prompt for write operations",
			Destination: &r.Yes,
		},
	}
}

func NewApp() *cli.Command {
	r := &Runtime{Output: output.DefaultOptions()}
	return &cli.Command{
		Name:  "tossinvest-cli",
		Usage: "Toss Invest Open API CLI",
		Flags: globalFlags(r),
		Before: func(ctx context.Context, _ *cli.Command) (context.Context, error) {
			if err := r.Init(); err != nil {
				return ctx, err
			}
			return ctx, nil
		},
		Commands: []*cli.Command{
			{
				Name:   "version",
				Usage:  "Show version",
				Action: func(_ context.Context, c *cli.Command) error {
					_, err := fmt.Fprintln(c.Root().Writer, Version)
					return err
				},
			},
			configCommand(r),
			authCommand(r),
			orderbooksCommand(r),
			pricesCommand(r),
			tradesCommand(r),
			priceLimitsCommand(r),
			candlesCommand(r),
			stocksCommand(r),
			exchangeRateCommand(r),
			marketCalendarCommand(r),
			accountsCommand(r),
			holdingsCommand(r),
			buyingPowerCommand(r),
			sellableQuantityCommand(r),
			commissionsCommand(r),
			ordersCommand(r),
		},
	}
}

func apiGet(r *Runtime, ctx context.Context, path string, query url.Values, needAccount bool) error {
	var seq *int64
	var err error
	if needAccount {
		seq, err = r.accountSeq(ctx)
		if err != nil {
			return err
		}
	}
	data, err := r.Client.GetJSON(ctx, path, query, seq)
	if err != nil {
		return err
	}
	return r.emit(data)
}

func apiPost(r *Runtime, ctx context.Context, path string, body any, needAccount bool) error {
	var seq *int64
	var err error
	if needAccount {
		seq, err = r.accountSeq(ctx)
		if err != nil {
			return err
		}
	}
	data, err := r.Client.PostJSON(ctx, path, body, seq)
	if err != nil {
		return err
	}
	return r.emit(data)
}
