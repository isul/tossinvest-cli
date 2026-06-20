# Price Limits

## get

Get daily upper/lower price limits for a symbol.

```bash
tossinvest-cli price-limits get --symbol 005930
```

### Flags

| Flag | Required | Description |
|---|---|---|
| `--symbol` | yes | Stock symbol |

### Response fields

| Field | Description |
|---|---|
| `upperLimit` | Upper price limit |
| `lowerLimit` | Lower price limit |
| `currency` | Price currency |
