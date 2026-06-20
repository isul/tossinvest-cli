# Prices

## list

Get current prices for one or more symbols (max 200, comma-separated).

```bash
tossinvest-cli prices list --symbols 005930
tossinvest-cli prices list --symbols 005930,AAPL,MSFT
```

### Flags

| Flag | Required | Description |
|---|---|---|
| `--symbols` | yes | Comma-separated symbols |

### Response fields

| Field | Description |
|---|---|
| `symbol` | Stock symbol |
| `price` | Current price |
| `change` | Price change |
| `changeRate` | Change rate |
| `currency` | KRW or USD |
