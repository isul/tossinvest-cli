# Order Info

Pre-trade checks: buying power, sellable quantity, commissions.

## buying-power get

Cash-based buying power (excludes margin).

```bash
tossinvest-cli buying-power get
```

## sellable-quantity get

Sellable quantity for a symbol.

```bash
tossinvest-cli sellable-quantity get --symbol 005930
tossinvest-cli sellable-quantity get --symbol AAPL
```

| Flag | Required |
|---|---|
| `--symbol` | yes |

## commissions list

Commission rates by market.

```bash
tossinvest-cli commissions list
```

Run these before `orders create` to validate available funds and fees.
