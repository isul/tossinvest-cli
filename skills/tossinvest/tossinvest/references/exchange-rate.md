# Exchange Rate

## get

Get KRW ↔ USD reference exchange rate (updated every minute).

```bash
tossinvest-cli exchange-rate get
tossinvest-cli exchange-rate get --date-time "2026-03-25T09:00:00+09:00"
```

### Flags

| Flag | Required | Description |
|---|---|---|
| `--date-time` | no | Reference datetime (ISO 8601). Omit for current rate. |

### Notes

- Reference/display rate only; actual order rate may differ
- `validFrom` / `validUntil` indicate the 1-minute validity window
