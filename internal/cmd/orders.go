package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/urfave/cli/v3"
)

func stocksCommand(r *Runtime) *cli.Command {
	return &cli.Command{
		Name:  "stocks",
		Usage: "Stock master data",
		Commands: []*cli.Command{
			{
				Name:  "list",
				Usage: "Get stock info (comma-separated, max 200)",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "symbols", Required: true, Usage: "Symbols e.g. 005930,AAPL"},
				},
				Action: func(ctx context.Context, c *cli.Command) error {
					q := url.Values{}
					q.Set("symbols", c.String("symbols"))
					return apiGet(r, ctx, "/api/v1/stocks", q, false)
				},
			},
			{
				Name:  "list-warnings",
				Usage: "Get buy warnings for a symbol",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "symbol", Required: true, Usage: "Stock symbol"},
				},
				Action: func(ctx context.Context, c *cli.Command) error {
					symbol := c.String("symbol")
					return apiGet(r, ctx, "/api/v1/stocks/"+url.PathEscape(symbol)+"/warnings", nil, false)
				},
			},
		},
	}
}

func exchangeRateCommand(r *Runtime) *cli.Command {
	return &cli.Command{
		Name:  "exchange-rate",
		Usage: "Exchange rate data",
		Commands: []*cli.Command{
			{
				Name:  "get",
				Usage: "Get KRW/USD exchange rate",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "date-time", Usage: "Reference datetime (ISO 8601)"},
				},
				Action: func(ctx context.Context, c *cli.Command) error {
					q := url.Values{}
					if v := c.String("date-time"); v != "" {
						q.Set("dateTime", v)
					}
					return apiGet(r, ctx, "/api/v1/exchange-rate", q, false)
				},
			},
		},
	}
}

func marketCalendarCommand(r *Runtime) *cli.Command {
	return &cli.Command{
		Name:  "market-calendar",
		Usage: "Market calendar",
		Commands: []*cli.Command{
			{
				Name:  "get-kr",
				Usage: "Get KR market calendar",
				Action: func(ctx context.Context, _ *cli.Command) error {
					return apiGet(r, ctx, "/api/v1/market-calendar/KR", nil, false)
				},
			},
			{
				Name:  "get-us",
				Usage: "Get US market calendar",
				Action: func(ctx context.Context, _ *cli.Command) error {
					return apiGet(r, ctx, "/api/v1/market-calendar/US", nil, false)
				},
			},
		},
	}
}

func accountsCommand(r *Runtime) *cli.Command {
	return &cli.Command{
		Name:  "accounts",
		Usage: "Account management",
		Commands: []*cli.Command{
			{
				Name:  "list",
				Usage: "List accounts",
				Action: func(ctx context.Context, _ *cli.Command) error {
					return apiGet(r, ctx, "/api/v1/accounts", nil, false)
				},
			},
		},
	}
}

func holdingsCommand(r *Runtime) *cli.Command {
	return &cli.Command{
		Name:  "holdings",
		Usage: "Holdings (보유 주식)",
		Commands: []*cli.Command{
			{
				Name:  "list",
				Usage: "List holdings",
				Action: func(ctx context.Context, _ *cli.Command) error {
					return apiGet(r, ctx, "/api/v1/holdings", nil, true)
				},
			},
		},
	}
}

func buyingPowerCommand(r *Runtime) *cli.Command {
	return &cli.Command{
		Name:  "buying-power",
		Usage: "Buying power (매수 가능 금액)",
		Commands: []*cli.Command{
			{
				Name:  "get",
				Usage: "Get buying power",
				Action: func(ctx context.Context, _ *cli.Command) error {
					return apiGet(r, ctx, "/api/v1/buying-power", nil, true)
				},
			},
		},
	}
}

func sellableQuantityCommand(r *Runtime) *cli.Command {
	return &cli.Command{
		Name:  "sellable-quantity",
		Usage: "Sellable quantity (판매 가능 수량)",
		Commands: []*cli.Command{
			{
				Name:  "get",
				Usage: "Get sellable quantity for a symbol",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "symbol", Required: true, Usage: "Stock symbol"},
				},
				Action: func(ctx context.Context, c *cli.Command) error {
					q := url.Values{}
					q.Set("symbol", c.String("symbol"))
					return apiGet(r, ctx, "/api/v1/sellable-quantity", q, true)
				},
			},
		},
	}
}

func commissionsCommand(r *Runtime) *cli.Command {
	return &cli.Command{
		Name:  "commissions",
		Usage: "Commission rates",
		Commands: []*cli.Command{
			{
				Name:  "list",
				Usage: "List commission rates",
				Action: func(ctx context.Context, _ *cli.Command) error {
					return apiGet(r, ctx, "/api/v1/commissions", nil, true)
				},
			},
		},
	}
}

func ordersCommand(r *Runtime) *cli.Command {
	return &cli.Command{
		Name:  "orders",
		Usage: "Order management",
		Commands: []*cli.Command{
			ordersListOpen(r),
			ordersListClosed(r),
			ordersRetrieve(r),
			ordersCreate(r),
			ordersModify(r),
			ordersCancel(r),
		},
	}
}

func ordersListOpen(r *Runtime) *cli.Command {
	return &cli.Command{
		Name:  "list-open",
		Usage: "List open orders",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "symbol", Usage: "Filter by symbol"},
			&cli.StringFlag{Name: "from", Usage: "Start date (YYYY-MM-DD, KST)"},
			&cli.StringFlag{Name: "to", Usage: "End date (YYYY-MM-DD, KST)"},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			q := url.Values{}
			q.Set("status", "OPEN")
			if v := c.String("symbol"); v != "" {
				q.Set("symbol", v)
			}
			if v := c.String("from"); v != "" {
				q.Set("from", v)
			}
			if v := c.String("to"); v != "" {
				q.Set("to", v)
			}
			return apiGet(r, ctx, "/api/v1/orders", q, true)
		},
	}
}

func ordersListClosed(r *Runtime) *cli.Command {
	return &cli.Command{
		Name:  "list-closed",
		Usage: "List closed orders",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "symbol", Usage: "Filter by symbol"},
			&cli.StringFlag{Name: "from", Usage: "Start date (YYYY-MM-DD, KST)"},
			&cli.StringFlag{Name: "to", Usage: "End date (YYYY-MM-DD, KST)"},
			&cli.IntFlag{Name: "limit", Usage: "Page size (max 100)", Value: 20},
			&cli.StringFlag{Name: "cursor", Usage: "Pagination cursor"},
			&cli.BoolFlag{Name: "all-pages", Usage: "Fetch all pages"},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			return fetchClosedOrders(r, ctx, c)
		},
	}
}

func fetchClosedOrders(r *Runtime, ctx context.Context, c *cli.Command) error {
	cursor := c.String("cursor")
	var allOrders []json.RawMessage

	for {
		q := url.Values{}
		q.Set("status", "CLOSED")
		if v := c.String("symbol"); v != "" {
			q.Set("symbol", v)
		}
		if v := c.String("from"); v != "" {
			q.Set("from", v)
		}
		if v := c.String("to"); v != "" {
			q.Set("to", v)
		}
		q.Set("limit", fmt.Sprintf("%d", c.Int("limit")))
		if cursor != "" {
			q.Set("cursor", cursor)
		}

		seq, err := r.accountSeq(ctx)
		if err != nil {
			return err
		}
		data, err := r.Client.GetJSON(ctx, "/api/v1/orders", q, seq)
		if err != nil {
			return err
		}

		var page struct {
			Orders     []json.RawMessage `json:"orders"`
			NextCursor *string           `json:"nextCursor"`
			HasNext    bool              `json:"hasNext"`
		}
		if err := json.Unmarshal(data, &page); err != nil {
			return r.emit(data)
		}

		allOrders = append(allOrders, page.Orders...)
		if !c.Bool("all-pages") || !page.HasNext || page.NextCursor == nil || *page.NextCursor == "" {
			if c.Bool("all-pages") {
				return r.emitJSON(map[string]any{
					"orders":     allOrders,
					"nextCursor": page.NextCursor,
					"hasNext":    false,
				})
			}
			return r.emit(data)
		}
		cursor = *page.NextCursor
	}
}

func ordersRetrieve(r *Runtime) *cli.Command {
	return &cli.Command{
		Name:  "retrieve",
		Usage: "Get order detail",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "order-id", Required: true, Usage: "Order ID"},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			id := c.String("order-id")
			return apiGet(r, ctx, "/api/v1/orders/"+url.PathEscape(id), nil, true)
		},
	}
}

func ordersCreate(r *Runtime) *cli.Command {
	return &cli.Command{
		Name:  "create",
		Usage: "Create an order",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "symbol", Required: true},
			&cli.StringFlag{Name: "side", Required: true, Usage: "BUY or SELL"},
			&cli.StringFlag{Name: "order-type", Required: true, Usage: "LIMIT or MARKET"},
			&cli.StringFlag{Name: "quantity", Usage: "Order quantity (integer shares)"},
			&cli.StringFlag{Name: "order-amount", Usage: "Order amount in USD (US MARKET only)"},
			&cli.StringFlag{Name: "price", Usage: "Limit price"},
			&cli.StringFlag{Name: "time-in-force", Usage: "DAY or CLS", Value: "DAY"},
			&cli.StringFlag{Name: "client-order-id", Usage: "Idempotency key"},
			&cli.BoolFlag{Name: "confirm-high-value-order", Usage: "Confirm order >= 100M KRW"},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			body := map[string]any{
				"symbol":    c.String("symbol"),
				"side":      c.String("side"),
				"orderType": c.String("order-type"),
			}
			if v := c.String("time-in-force"); v != "" {
				body["timeInForce"] = v
			}
			if v := c.String("client-order-id"); v != "" {
				body["clientOrderId"] = v
			}
			if c.Bool("confirm-high-value-order") {
				body["confirmHighValueOrder"] = true
			}
			if v := c.String("quantity"); v != "" {
				body["quantity"] = v
			}
			if v := c.String("order-amount"); v != "" {
				body["orderAmount"] = v
			}
			if v := c.String("price"); v != "" {
				body["price"] = v
			}

			desc := fmt.Sprintf("orders create %s %s %s", c.String("side"), c.String("symbol"), c.String("order-type"))
			if err := confirmRequired(r, desc); err != nil {
				return err
			}
			return apiPost(r, ctx, "/api/v1/orders", body, true)
		},
	}
}

func ordersModify(r *Runtime) *cli.Command {
	return &cli.Command{
		Name:  "modify",
		Usage: "Modify an order",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "order-id", Required: true},
			&cli.StringFlag{Name: "order-type", Required: true, Usage: "LIMIT or MARKET"},
			&cli.StringFlag{Name: "quantity", Usage: "New quantity (required for KR)"},
			&cli.StringFlag{Name: "price", Usage: "New limit price"},
			&cli.BoolFlag{Name: "confirm-high-value-order", Usage: "Confirm order >= 100M KRW"},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			body := map[string]any{
				"orderType": c.String("order-type"),
			}
			if v := c.String("quantity"); v != "" {
				body["quantity"] = v
			}
			if v := c.String("price"); v != "" {
				body["price"] = v
			}
			if c.Bool("confirm-high-value-order") {
				body["confirmHighValueOrder"] = true
			}

			desc := fmt.Sprintf("orders modify --order-id %s", c.String("order-id"))
			if err := confirmRequired(r, desc); err != nil {
				return err
			}
			id := c.String("order-id")
			return apiPost(r, ctx, "/api/v1/orders/"+url.PathEscape(id)+"/modify", body, true)
		},
	}
}

func ordersCancel(r *Runtime) *cli.Command {
	return &cli.Command{
		Name:  "cancel",
		Usage: "Cancel an order",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "order-id", Required: true, Usage: "Order ID"},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			desc := fmt.Sprintf("orders cancel --order-id %s", c.String("order-id"))
			if err := confirmRequired(r, desc); err != nil {
				return err
			}
			id := c.String("order-id")
			return apiPost(r, ctx, "/api/v1/orders/"+url.PathEscape(id)+"/cancel", map[string]any{}, true)
		},
	}
}

func confirmRequired(r *Runtime, desc string) error {
	return confirmRequiredYes(r.Yes, desc)
}
