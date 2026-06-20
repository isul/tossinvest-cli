# Orderbooks

## list

Get bid/ask orderbook for a symbol.

```bash
tossinvest-cli orderbooks list --symbol 005930
tossinvest-cli orderbooks list --symbol AAPL
```

### Flags

| Flag | Required | Description |
|---|---|---|
| `--symbol` | yes | Stock symbol (KR 6-digit or US ticker) |

### Response fields

| Field | Description |
|---|---|
| `asks` | Sell-side quotes (price, volume) |
| `bids` | Buy-side quotes (price, volume) |
| `timestamp` | Quote timestamp |
| `currency` | Quote currency (KRW/USD) |
