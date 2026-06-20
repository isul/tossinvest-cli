# Orders

## list-open

List in-progress orders (returns all, no pagination).

```bash
tossinvest-cli orders list-open
tossinvest-cli orders list-open --symbol 005930
tossinvest-cli orders list-open --from 2026-03-01 --to 2026-03-31
```

| Flag | Description |
|---|---|
| `--symbol` | Filter by symbol |
| `--from` | Start date (YYYY-MM-DD, KST) |
| `--to` | End date (YYYY-MM-DD, KST) |

## list-closed

List completed/terminated orders (paginated).

```bash
tossinvest-cli orders list-closed --limit 20
tossinvest-cli orders list-closed --cursor "<nextCursor>"
tossinvest-cli orders list-closed --all-pages
```

| Flag | Default | Description |
|---|---|---|
| `--limit` | 20 | Page size (max 100) |
| `--cursor` | | Pagination cursor from prior response |
| `--all-pages` | false | Fetch all pages |

## retrieve

Get order detail by ID.

```bash
tossinvest-cli orders retrieve --order-id "<orderId>"
```

## create

**Requires CONFIRM** before execution.

### KR limit buy

```bash
tossinvest-cli orders create \
  --symbol 005930 \
  --side BUY \
  --order-type LIMIT \
  --quantity 10 \
  --price 70000 \
  --client-order-id my-order-001
```

### US market buy (USD amount)

```bash
tossinvest-cli orders create \
  --symbol AAPL \
  --side BUY \
  --order-type MARKET \
  --order-amount 100.5
```

### US LOC (LIMIT + CLS)

```bash
tossinvest-cli orders create \
  --symbol AAPL \
  --side BUY \
  --order-type LIMIT \
  --time-in-force CLS \
  --quantity 10 \
  --price 185.5
```

| Flag | Required | Description |
|---|---|---|
| `--symbol` | yes | Stock symbol |
| `--side` | yes | `BUY` or `SELL` |
| `--order-type` | yes | `LIMIT` or `MARKET` |
| `--quantity` | one of qty/amount | Integer shares |
| `--order-amount` | one of qty/amount | USD amount (US MARKET, MARKET only) |
| `--price` | LIMIT only | Limit price |
| `--time-in-force` | no | `DAY` (default) or `CLS` |
| `--client-order-id` | no | Idempotency key (10 min TTL) |
| `--confirm-high-value-order` | 100M+ KRW | Required for orders >= 100M KRW |

## modify

**Requires CONFIRM.** KR requires `--quantity`; US price-only (no quantity).

```bash
tossinvest-cli orders modify \
  --order-id "<orderId>" \
  --order-type LIMIT \
  --quantity 15 \
  --price 71000
```

## cancel

**Requires CONFIRM.**

```bash
tossinvest-cli orders cancel --order-id "<orderId>"
```

## Rules

- `quantity` and `orderAmount` are mutually exclusive
- `orderAmount` only for US MARKET + MARKET, regular hours only
- KR prices must match tick size; US: 4 decimals under $1, 2 decimals otherwise
- Use `--yes` only for automation (skips CONFIRM prompt)
