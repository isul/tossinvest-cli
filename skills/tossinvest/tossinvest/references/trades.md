# Trades

## list

Get recent trades for a symbol (today).

```bash
tossinvest-cli trades list --symbol 005930
tossinvest-cli trades list --symbol AAPL
```

### Flags

| Flag | Required | Description |
|---|---|---|
| `--symbol` | yes | Stock symbol |

### Response fields

| Field | Description |
|---|---|
| `price` | Trade price |
| `volume` | Trade volume |
| `timestamp` | Trade time |
| `side` | Aggressor side if available |
