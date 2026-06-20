# Stocks

## list

Get stock master data (max 200 symbols, comma-separated).

```bash
tossinvest-cli stocks list --symbols 005930
tossinvest-cli stocks list --symbols 005930,AAPL
```

### Flags

| Flag | Required | Description |
|---|---|---|
| `--symbols` | yes | Comma-separated symbols |

### Key fields

| Field | Description |
|---|---|
| `name` | Korean name |
| `englishName` | English name |
| `market` | Market (KOSPI, NASDAQ, etc.) |
| `currency` | Trading currency |
| `status` | Listing status |
| `koreanMarketDetail` | KR-specific trading flags |

## list-warnings

Get buy warnings and VI (volatility interruption) info.

```bash
tossinvest-cli stocks list-warnings --symbol 005930
```

### Flags

| Flag | Required | Description |
|---|---|---|
| `--symbol` | yes | Stock symbol |

Warning types include: `LIQUIDATION_TRADING`, `OVERHEATED`, `INVESTMENT_WARNING`, `INVESTMENT_RISK`, `VI_STATIC`, `VI_DYNAMIC`, `STOCK_WARRANTS`

Empty array `[]` means no active warnings (not an error).
