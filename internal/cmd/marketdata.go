package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/urfave/cli/v3"
)

func orderbooksCommand(r *Runtime) *cli.Command {
	return &cli.Command{
		Name:  "orderbooks",
		Usage: "Orderbook (호가) data",
		Commands: []*cli.Command{
			{
				Name:  "list",
				Usage: "Get orderbook for a symbol",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "symbol", Required: true, Usage: "Stock symbol (005930, AAPL)"},
				},
				Action: func(ctx context.Context, c *cli.Command) error {
					q := url.Values{}
					q.Set("symbol", c.String("symbol"))
					return apiGet(r, ctx, "/api/v1/orderbook", q, false)
				},
			},
		},
	}
}

func pricesCommand(r *Runtime) *cli.Command {
	return &cli.Command{
		Name:  "prices",
		Usage: "Current price data",
		Commands: []*cli.Command{
			{
				Name:  "list",
				Usage: "Get current prices (comma-separated, max 200)",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "symbols", Required: true, Usage: "Symbols e.g. 005930,AAPL"},
				},
				Action: func(ctx context.Context, c *cli.Command) error {
					q := url.Values{}
					q.Set("symbols", c.String("symbols"))
					return apiGet(r, ctx, "/api/v1/prices", q, false)
				},
			},
		},
	}
}

func tradesCommand(r *Runtime) *cli.Command {
	return &cli.Command{
		Name:  "trades",
		Usage: "Recent trades",
		Commands: []*cli.Command{
			{
				Name:  "list",
				Usage: "Get recent trades for a symbol",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "symbol", Required: true, Usage: "Stock symbol"},
				},
				Action: func(ctx context.Context, c *cli.Command) error {
					q := url.Values{}
					q.Set("symbol", c.String("symbol"))
					return apiGet(r, ctx, "/api/v1/trades", q, false)
				},
			},
		},
	}
}

func priceLimitsCommand(r *Runtime) *cli.Command {
	return &cli.Command{
		Name:  "price-limits",
		Usage: "Daily price limits",
		Commands: []*cli.Command{
			{
				Name:  "get",
				Usage: "Get upper/lower price limits",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "symbol", Required: true, Usage: "Stock symbol"},
				},
				Action: func(ctx context.Context, c *cli.Command) error {
					q := url.Values{}
					q.Set("symbol", c.String("symbol"))
					return apiGet(r, ctx, "/api/v1/price-limits", q, false)
				},
			},
		},
	}
}

func candlesCommand(r *Runtime) *cli.Command {
	return &cli.Command{
		Name:  "candles",
		Usage: "Candle chart data",
		Commands: []*cli.Command{
			{
				Name:  "list",
				Usage: "Get candle data",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "symbol", Required: true, Usage: "Stock symbol"},
					&cli.StringFlag{Name: "interval", Required: true, Usage: "Candle interval: 1m or 1d"},
					&cli.IntFlag{Name: "count", Usage: "Number of candles (max 200)", Value: 100},
					&cli.StringFlag{Name: "before", Usage: "Pagination upper bound (ISO 8601)"},
					&cli.BoolFlag{Name: "adjusted", Usage: "Use adjusted prices", Value: true},
					&cli.BoolFlag{Name: "all-pages", Usage: "Fetch all pages"},
				},
				Action: func(ctx context.Context, c *cli.Command) error {
					return fetchCandles(r, ctx, c)
				},
			},
		},
	}
}

func fetchCandles(r *Runtime, ctx context.Context, c *cli.Command) error {
	var allCandles []json.RawMessage
	before := c.String("before")

	for {
		q := url.Values{}
		q.Set("symbol", c.String("symbol"))
		q.Set("interval", c.String("interval"))
		q.Set("count", fmt.Sprintf("%d", c.Int("count")))
		if before != "" {
			q.Set("before", before)
		}
		q.Set("adjusted", fmt.Sprintf("%t", c.Bool("adjusted")))

		data, err := r.Client.GetJSON(ctx, "/api/v1/candles", q, nil)
		if err != nil {
			return err
		}

		var page struct {
			Candles    []json.RawMessage `json:"candles"`
			NextBefore *string           `json:"nextBefore"`
		}
		if err := json.Unmarshal(data, &page); err != nil {
			return r.emit(data)
		}

		allCandles = append(allCandles, page.Candles...)
		if !c.Bool("all-pages") || page.NextBefore == nil || *page.NextBefore == "" {
			if c.Bool("all-pages") && len(allCandles) > 0 {
				return r.emitJSON(map[string]any{"candles": allCandles})
			}
			return r.emit(data)
		}
		before = *page.NextBefore
	}
}
