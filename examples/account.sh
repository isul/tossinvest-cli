#!/usr/bin/env bash
# Account-scoped examples (requires credentials + account)
set -euo pipefail

echo "=== Accounts ==="
tossinvest-cli accounts list --format pretty

echo "=== Holdings ==="
tossinvest-cli holdings list --format pretty

echo "=== Buying power ==="
tossinvest-cli buying-power get --format pretty

echo "=== Open orders ==="
tossinvest-cli orders list-open --format pretty
