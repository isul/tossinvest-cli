[한국어](./README_KR.md) | English

# tossinvest-cli

Official-style CLI for the [Toss Invest Open API](https://developers.tossinvest.com/docs).

> **This program is not an official Toss Invest product.**

## Install

### npm

```sh
npm install -g @isul/tossinvest-cli
```

### Go

Requires Go 1.22+.

```sh
go install github.com/isul/tossinvest-cli/cmd/tossinvest-cli@latest
```

The binary is installed to `$GOPATH/bin` (default `$HOME/go/bin`). Add it to your PATH if needed.

### Local development

```sh
git clone https://github.com/isul/tossinvest-cli.git
cd tossinvest-cli
./scripts/run --help
```

## Usage

Resource-based command structure:

```sh
tossinvest-cli [resource] <command> [flags...]
```

```sh
tossinvest-cli config set
tossinvest-cli prices list --symbols 005930,AAPL
tossinvest-cli holdings list
tossinvest-cli orders create --symbol 005930 --side BUY --order-type LIMIT --quantity 10 --price 70000
```

See `tossinvest-cli <resource> <command> --help` for details. Examples are in [`examples/`](examples/).

## Environment variables

| Variable | Description | Required |
| --- | --- | --- |
| `TOSSINVEST_CLIENT_ID` | OAuth client ID | No |
| `TOSSINVEST_CLIENT_SECRET` | OAuth client secret | No |
| `TOSSINVEST_ACCOUNT_SEQ` | Default account sequence | No |
| `TOSSINVEST_BASE_URL` | API base URL | No (default: `https://openapi.tossinvest.com`) |
| `TOSSINVEST_AUTO_CONFIRM` | Skip CONFIRM for write ops when `1` | No |

## Global flags

- `--client-id` — OAuth client ID (`TOSSINVEST_CLIENT_ID`)
- `--client-secret` — OAuth client secret (`TOSSINVEST_CLIENT_SECRET`)
- `--account-seq` — Account header value (`TOSSINVEST_ACCOUNT_SEQ`)
- `--base-url` — Custom API URL
- `--format` — Output format: `auto`, `json`, `yaml`, `pretty`, `raw`
- `--format-error` — Error output format
- `--transform` — [GJSON](https://github.com/tidwall/gjson) output transform
- `--debug` — HTTP debug logging
- `--yes` — Skip CONFIRM for write operations

## AI Agent Skill

Agent skill for Cursor and other AI tools is in [`skills/tossinvest/tossinvest/SKILL.md`](skills/tossinvest/tossinvest/SKILL.md).

```sh
mkdir -p ~/.cursor/skills
cp -r skills/tossinvest/tossinvest ~/.cursor/skills/tossinvest
```

## API coverage

All endpoints from [OpenAPI v1.1.1](https://openapi.tossinvest.com/openapi-docs/latest/openapi.json):

- Auth (OAuth2 token, internal)
- Market data: orderbook, prices, trades, price limits, candles
- Stock info: stocks, warnings
- Market info: exchange rate, KR/US calendar
- Account & asset: accounts, holdings
- Orders: create, modify, cancel, list, detail
- Order info: buying power, sellable quantity, commissions

## License

MIT — see [LICENSE](LICENSE).
