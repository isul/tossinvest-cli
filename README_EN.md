[한국어](./README.md) | English

# tossinvest-cli — Command-line (CLI) tool for the Toss Invest Open API

A command-line (CLI) tool for the [Toss Invest Open API](https://developers.tossinvest.com/docs). Quickly **look up domestic and overseas stock prices**, **check Toss Invest account balances**, and **place stock orders / trade** right from your terminal — ideal for scripting, automation, and AI coding agent integration.

> **This program is not an official Toss Invest product.**

**Keywords**: Toss Invest, 토스증권, Toss Invest API, stock CLI, stock quotes, account balance, stock order, automated trading, Go, open source

## Features

- **Domestic & overseas stock price lookup** — real-time orderbook, current price, trades, price limits, and candles (minute/daily)
- **Stock info** — stock fundamentals and buy warnings
- **Toss Invest account balance** — list accounts and view holdings (balances)
- **CLI-based stock ordering / trading** — create, modify, and cancel buy/sell orders; list and inspect orders
- **Order info** — buying power, sellable quantity, and commission calculation
- **Market info** — exchange rate and KR/US market open calendars
- **Flexible output formats** — `json`, `yaml`, `pretty`, `raw`, and [GJSON](https://github.com/tidwall/gjson) transforms, optimized for scripting
- **AI agent integration** — ships a `SKILL.md` standard skill for Claude Code, Codex, Copilot, Gemini CLI, OpenClaw, Hermes, and more
- **Cross-platform** — prebuilt binaries for Linux, macOS, and Windows (amd64/arm64)

## Quick Start

### 1. Install (per OS)

Grab the latest prebuilt binary from GitHub Releases. Use the command that matches your OS/architecture.

**Linux (amd64)**

```sh
curl -fsSL https://github.com/isul/tossinvest-cli/releases/latest/download/tossinvest-cli_linux_amd64.tar.gz \
  | tar -xz tossinvest-cli && sudo mv tossinvest-cli /usr/local/bin/
```

**Linux (arm64)**

```sh
curl -fsSL https://github.com/isul/tossinvest-cli/releases/latest/download/tossinvest-cli_linux_arm64.tar.gz \
  | tar -xz tossinvest-cli && sudo mv tossinvest-cli /usr/local/bin/
```

**macOS (Apple Silicon / arm64)**

```sh
curl -fsSL https://github.com/isul/tossinvest-cli/releases/latest/download/tossinvest-cli_darwin_arm64.tar.gz \
  | tar -xz tossinvest-cli && sudo mv tossinvest-cli /usr/local/bin/
```

**macOS (Intel / amd64)**

```sh
curl -fsSL https://github.com/isul/tossinvest-cli/releases/latest/download/tossinvest-cli_darwin_amd64.tar.gz \
  | tar -xz tossinvest-cli && sudo mv tossinvest-cli /usr/local/bin/
```

**Windows (PowerShell, amd64)**

```powershell
Invoke-WebRequest -Uri "https://github.com/isul/tossinvest-cli/releases/latest/download/tossinvest-cli_windows_amd64.zip" -OutFile tossinvest-cli.zip
Expand-Archive tossinvest-cli.zip -DestinationPath .
.\tossinvest-cli.exe version
```

> Alternative: if you have Go installed, run `go install github.com/isul/tossinvest-cli/cmd/tossinvest-cli@latest`.

### 2. Configure & use

```sh
# Configure credentials (OAuth client ID/secret)
tossinvest-cli config set

# Look up domestic & overseas stock prices
tossinvest-cli prices list --symbols 005930,AAPL

# Check account balance (holdings)
tossinvest-cli holdings list

# Place a stock order (buy)
tossinvest-cli orders create --symbol 005930 --side BUY --order-type LIMIT --quantity 10 --price 70000
```

For full installation options, see the [Install](#install) section below; for per-command flags, use `--help`.

## Install

### GitHub Releases

Pre-built binaries for Linux, macOS, and Windows (amd64 and arm64) are available on [GitHub Releases](https://github.com/isul/tossinvest-cli/releases).

1. Open the [latest release](https://github.com/isul/tossinvest-cli/releases/latest)
2. Download the archive for your OS and architecture (e.g. `tossinvest-cli_linux_amd64.tar.gz`)
3. Extract the binary and run it (add to your PATH for global use)

Linux and macOS use `.tar.gz`; Windows uses `.zip`.

```sh
tar -xzf tossinvest-cli_linux_amd64.tar.gz
sudo mv tossinvest-cli /usr/local/bin/
tossinvest-cli version
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

An [Agent Skills](https://agentskills.io/) skill for AI coding agents is in [`skills/tossinvest/tossinvest/SKILL.md`](skills/tossinvest/tossinvest/SKILL.md). It follows the open `SKILL.md` format and works with Claude Code, Codex, Copilot, Gemini CLI, OpenClaw, Hermes, and other compatible agents.

Copy the skill directory into your agent's skills folder:

| Agent | Skills directory |
| --- | --- |
| Claude Code | `~/.claude/skills/` or `.claude/skills/` |
| OpenAI Codex | `~/.agents/skills/` or `.agents/skills/` |
| GitHub Copilot | `.github/skills/` |
| Cursor | `~/.cursor/skills/` |

```sh
cp -r skills/tossinvest/tossinvest ~/.claude/skills/tossinvest
```

Adjust the destination path for your agent.

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
