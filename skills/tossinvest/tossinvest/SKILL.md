---
name: tossinvest
description: >
  Use tossinvest-cli for Toss Invest Open API — stock prices, orderbooks, candles, accounts, holdings, orders.
  토스증권 CLI로 시세 조회, 주문, 잔고 확인, 매수/매도를 처리합니다.
  Trigger this skill whenever the user wants to query stock prices, place/cancel/modify orders, check holdings or buying power, or interact with the Toss Invest Open API — in any language.
  사용자가 토스증권 시세·주문·잔고·매매를 언급하면 반드시 이 스킬을 사용하세요.
metadata:
  version: v0.1.0
  author: isul
license: MIT
---

# TossInvest Skill

Use the `tossinvest-cli` binary for all Toss Invest Open API interactions.

## Language Behavior

Detect the user's language and respond accordingly:

- **Korean user**: respond in Korean, use Korean terminology from `references/glossary.md`
- **English user**: respond in English, use English terminology from the same glossary
- **Mixed/ambiguous**: follow the language of the most recent message

Load `references/glossary.md` when translating terminology or explaining response fields.

## Setup

If `tossinvest-cli` is not installed or credentials are not configured, load `references/setup.md` and follow the steps there.

Check if `tossinvest-cli` is available:

```bash
tossinvest-cli version
```

## Authentication

Private endpoints require OAuth 2.0 Client Credentials. Configure via the CLI (recommended):

```bash
tossinvest-cli config set
```

Credentials are saved to `~/.tossinvest/config.yaml` and automatically used for all CLI commands.

Alternatively, set via environment variables:

```bash
export TOSSINVEST_CLIENT_ID=<your-client-id>
export TOSSINVEST_CLIENT_SECRET=<your-client-secret>
```

Or pass inline per command:

```bash
tossinvest-cli --client-id <id> --client-secret <secret> <resource> <command> ...
```

**Private** (require auth + account header): `accounts`, `holdings`, `buying-power`, `sellable-quantity`, `commissions`, `orders`
**Public** (token only): `orderbooks`, `prices`, `trades`, `price-limits`, `candles`, `stocks`, `exchange-rate`, `market-calendar`

## Account Header

Account-scoped APIs require `X-Tossinvest-Account` (account sequence from `accounts list`):

```bash
tossinvest-cli accounts list
tossinvest-cli --account-seq 1 holdings list
```

Set default account via `config set` or `TOSSINVEST_ACCOUNT_SEQ`. If exactly one account exists, the CLI auto-selects it.

## Safety Rule — Write Operations

Before executing any write operation, show the full command and ask the user to type `CONFIRM`.

Write operations:
- `orders create`, `orders modify`, `orders cancel`

Use `--yes` or `TOSSINVEST_AUTO_CONFIRM=1` only for scripted/automation scenarios (document the risk).

## Symbol Format

| Market | Format | Example |
|---|---|---|
| KR (KRX) | 6-digit code | `005930` (Samsung Electronics) |
| US | Ticker | `AAPL` |

## Order Domain Concepts

### Side

| Value | Meaning |
|---|---|
| `BUY` | Buy (매수) |
| `SELL` | Sell (매도) |

### Order Type

| Value | Meaning |
|---|---|
| `LIMIT` | Limit order (지정가) |
| `MARKET` | Market order (시장가) |

### Time In Force

| Value | Meaning |
|---|---|
| `DAY` | Valid until end of regular session (default) |
| `CLS` | At-the-close (US LIMIT only, LOC) |

### Quantity vs Amount

- `quantity`: integer shares — use for KR and US quantity-based orders
- `orderAmount`: USD amount — US MARKET only, `MARKET` orders, regular hours only

### Order Status (list filter vs detail)

List filter groups:
- `OPEN`: PENDING, PARTIAL_FILLED, PENDING_CANCEL, PENDING_REPLACE
- `CLOSED`: FILLED, CANCELED, REJECTED, REPLACED, etc.

### Before First Order

Run these before placing orders on an unfamiliar symbol:

```bash
tossinvest-cli buying-power get
tossinvest-cli sellable-quantity get --symbol AAPL
tossinvest-cli commissions list
tossinvest-cli stocks list-warnings --symbol 005930
```

## Command Reference

When you need detailed flag information for a resource, read the corresponding reference file.

| Resource | Subcommands | Reference |
|---|---|---|
| `orders` | create, modify, cancel, list-open, list-closed, retrieve | [`references/orders.md`](references/orders.md) |
| `orderbooks` | list | [`references/orderbooks.md`](references/orderbooks.md) |
| `prices` | list | [`references/prices.md`](references/prices.md) |
| `candles` | list | [`references/candles.md`](references/candles.md) |
| `trades` | list | [`references/trades.md`](references/trades.md) |
| `price-limits` | get | [`references/price-limits.md`](references/price-limits.md) |
| `stocks` | list, list-warnings | [`references/stocks.md`](references/stocks.md) |
| `exchange-rate` | get | [`references/exchange-rate.md`](references/exchange-rate.md) |
| `market-calendar` | get-kr, get-us | [`references/market-calendar.md`](references/market-calendar.md) |
| `accounts` / `holdings` | list | [`references/accounts.md`](references/accounts.md), [`references/holdings.md`](references/holdings.md) |
| `buying-power` / `sellable-quantity` / `commissions` | get / get / list | [`references/order-info.md`](references/order-info.md) |
| Output & Filtering | --format, --transform | [`references/output.md`](references/output.md) |
| Glossary | Term translations | [`references/glossary.md`](references/glossary.md) |
| Setup | Installation, credentials | [`references/setup.md`](references/setup.md) |

For flags not listed in reference files, run: `tossinvest-cli <resource> <command> --help`

## Environment

```bash
tossinvest-cli accounts list
tossinvest-cli --base-url https://openapi.tossinvest.com prices list --symbols 005930
```
