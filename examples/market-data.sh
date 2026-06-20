#!/usr/bin/env bash
# Public market data examples (requires valid credentials)
set -euo pipefail

echo "=== Samsung Electronics price ==="
tossinvest-cli prices list --symbols 005930 --format pretty

echo "=== KR market calendar ==="
tossinvest-cli market-calendar get-kr --format pretty

echo "=== Exchange rate ==="
tossinvest-cli exchange-rate get --format pretty
