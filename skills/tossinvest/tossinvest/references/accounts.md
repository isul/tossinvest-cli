# Accounts

## list

List brokerage accounts. Call this first to obtain `accountSeq`.

```bash
tossinvest-cli accounts list
```

### Response fields

| Field | Description |
|---|---|
| `accountSeq` | Account sequence — use as `--account-seq` or in config |
| `accountType` | Currently `BROKERAGE` only |
| `accountName` | Display name |

### Usage

```bash
# Save default account
tossinvest-cli config set   # enter account_seq when prompted

# Or per-command
tossinvest-cli --account-seq 1 holdings list
```

If exactly one account exists, account-scoped commands auto-select it.
