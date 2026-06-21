# Setup

## Install

### Go

Requires Go 1.22+.

```bash
go install github.com/isul/tossinvest-cli/cmd/tossinvest-cli@latest
```

Binary location: `$GOPATH/bin` or `$HOME/go/bin`. Add to PATH if needed.

### npm

```bash
npm install -g @isul/tossinvest-cli
```

### Local development

```bash
git clone https://github.com/isul/tossinvest-cli.git
cd tossinvest-cli
./scripts/run --help
```

## Verify

```bash
tossinvest-cli version
```

## API Credentials

1. Register at [Toss Invest Open API](https://developers.tossinvest.com/docs)
2. Obtain `client_id` and `client_secret` (OAuth 2.0 Client Credentials)
3. Configure:

```bash
tossinvest-cli config set
```

Config file: `~/.tossinvest/config.yaml`

## Environment Variables

| Variable | Description |
|---|---|
| `TOSSINVEST_CLIENT_ID` | OAuth client ID |
| `TOSSINVEST_CLIENT_SECRET` | OAuth client secret |
| `TOSSINVEST_ACCOUNT_SEQ` | Default account sequence |
| `TOSSINVEST_BASE_URL` | API base URL (default: `https://openapi.tossinvest.com`) |
| `TOSSINVEST_AUTO_CONFIRM` | Set to `1` to skip CONFIRM for write ops |

## Agent Skill

This repository includes an [Agent Skills](https://agentskills.io/) skill. Copy or symlink `skills/tossinvest/tossinvest` into your agent's skills folder:

| Agent | Skills directory |
|---|---|
| Claude Code | `~/.claude/skills/` or `.claude/skills/` |
| OpenAI Codex | `~/.agents/skills/` or `.agents/skills/` |
| GitHub Copilot | `.github/skills/` |
| Cursor | `~/.cursor/skills/` |

```bash
cp -r skills/tossinvest/tossinvest ~/.claude/skills/tossinvest
```

## Windows 11

```powershell
go install github.com/isul/tossinvest-cli/cmd/tossinvest-cli@latest
$env:Path += ";$(go env GOPATH)\bin"
tossinvest-cli version
```
