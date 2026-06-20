# Holdings

## list

Get holdings (보유 주식) with P&L summary. Requires account header.

```bash
tossinvest-cli holdings list
tossinvest-cli --account-seq 1 holdings list
```

### Response structure

| Field | Description |
|---|---|
| `overview` | Account-level summary (total value, profit/loss) |
| `items` | Per-symbol holdings |

### Item fields

| Field | Description |
|---|---|
| `symbol` | Stock symbol |
| `quantity` | Held quantity |
| `averagePurchasePrice` | Average cost |
| `currentPrice` | Current price |
| `evaluatedAmount` | Market value |
| `profitLoss` | Unrealized P&L (KRW converted) |

Covers KR and US stocks only (no options/bonds).
