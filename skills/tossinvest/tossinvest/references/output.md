# Output & Filtering

## Format

```bash
tossinvest-cli --format json prices list --symbols 005930
tossinvest-cli --format yaml holdings list
tossinvest-cli --format pretty accounts list
```

Supported: `auto`, `json`, `yaml`, `pretty`, `raw`

## Transform (GJSON)

Extract nested fields:

```bash
tossinvest-cli --format json --transform "0.symbol" stocks list --symbols 005930,AAPL
tossinvest-cli --format json --transform "orders.#.orderId" orders list-open
```

## Auto-paging

```bash
tossinvest-cli orders list-closed --all-pages
tossinvest-cli candles list --symbol 005930 --interval 1d --all-pages
```

## Debug

```bash
tossinvest-cli --debug auth token
```
