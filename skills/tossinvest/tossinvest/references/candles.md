# Candles

## list

Get OHLCV candle chart data.

```bash
tossinvest-cli candles list --symbol 005930 --interval 1d --count 100
tossinvest-cli candles list --symbol AAPL --interval 1m --count 60
tossinvest-cli candles list --symbol 005930 --interval 1d --all-pages
```

### Flags

| Flag | Required | Default | Description |
|---|---|---|---|
| `--symbol` | yes | | Stock symbol |
| `--interval` | yes | | `1m` (1-minute) or `1d` (daily) |
| `--count` | no | 100 | Number of candles (max 200) |
| `--before` | no | | Pagination cursor (ISO 8601, use `nextBefore` from prior response) |
| `--adjusted` | no | true | Apply adjusted prices |
| `--all-pages` | no | false | Fetch all pages |

### Response fields

| Field | Description |
|---|---|
| `candles` | Array of OHLCV bars |
| `nextBefore` | Cursor for next page (null if last page) |

Each candle: `openPrice`, `highPrice`, `lowPrice`, `closePrice`, `volume`, `timestamp`, `currency`
