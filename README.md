[English](./README_EN.md) | 한국어

# tossinvest-cli — 토스증권 Open API 커맨드라인(CLI) 도구

[토스증권 Open API](https://developers.tossinvest.com/docs)를 위한 커맨드라인(CLI) 도구입니다. 터미널에서 **국내·해외 주식 시세 조회**, **토스증권 계좌 잔고 확인**, **주식 주문/매매**를 빠르게 실행할 수 있으며, 스크립트와 자동화, AI 코딩 에이전트 연동에도 적합합니다.

> **이 프로그램은 토스증권 공식 프로그램이 아닙니다.**

**키워드**: 토스증권, Toss Invest, 토스증권 API, 주식 CLI, 주식 시세 조회, 계좌 잔고, 주식 주문, 자동매매, Go, 오픈소스

## 주요 기능

- **국내·해외 주식 시세 조회** — 호가, 현재가, 체결, 상하한가, 캔들(분/일봉) 등 실시간 시세 데이터 조회
- **종목 정보 조회** — 종목 기본정보 및 매수 유의사항 확인
- **토스증권 계좌 잔고 확인** — 계좌 목록 및 보유 주식(잔고) 조회
- **CLI 기반 주식 주문/매매** — 매수/매도 주문 생성, 정정, 취소, 주문 목록·상세 조회
- **주문 정보 확인** — 매수 가능 금액, 판매 가능 수량, 수수료 계산
- **시장 정보** — 환율, 국내/해외 장 개장 캘린더 조회
- **다양한 출력 형식** — `json`, `yaml`, `pretty`, `raw` 및 [GJSON](https://github.com/tidwall/gjson) 변환으로 스크립트 자동화에 최적화
- **AI 에이전트 연동** — `SKILL.md` 표준 스킬로 Claude Code, Codex, Copilot, Gemini CLI, OpenClaw, Hermes 등에서 사용 가능
- **크로스 플랫폼** — Linux, macOS, Windows(amd64/arm64) 바이너리 제공

## Quick Start

### 1. 설치 (OS별)

GitHub Releases의 최신 바이너리를 받습니다. 자신의 OS/아키텍처에 맞는 명령을 사용하세요.

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

> 대안: Go가 설치되어 있으면 `go install github.com/isul/tossinvest-cli/cmd/tossinvest-cli@latest`로도 설치할 수 있습니다.

### 2. 설정 및 사용

```sh
# 인증 정보 설정 (OAuth client ID/secret)
tossinvest-cli config set

# 국내·해외 주식 시세 조회
tossinvest-cli prices list --symbols 005930,AAPL

# 계좌 잔고(보유 주식) 확인
tossinvest-cli holdings list

# 주식 주문(매수) 실행
tossinvest-cli orders create --symbol 005930 --side BUY --order-type LIMIT --quantity 10 --price 70000
```

자세한 설치 방법은 아래 [설치](#설치) 섹션을, 명령별 옵션은 `--help`를 참고하세요.

## 설치

### GitHub Releases에서 바이너리 다운로드

Linux, macOS, Windows(amd64, arm64)용 빌드 바이너리는 [GitHub Releases](https://github.com/isul/tossinvest-cli/releases)에서 받을 수 있습니다.

1. [최신 릴리스](https://github.com/isul/tossinvest-cli/releases/latest) 페이지를 엽니다
2. 사용 중인 OS/아키텍처에 맞는 파일을 다운로드합니다 (예: `tossinvest-cli_linux_amd64.tar.gz`)
3. 압축을 풀고 바이너리를 실행합니다 (전역 사용을 원하면 PATH에 추가)

Linux/macOS는 `.tar.gz`, Windows는 `.zip` 형식입니다.

```sh
tar -xzf tossinvest-cli_linux_amd64.tar.gz
sudo mv tossinvest-cli /usr/local/bin/
tossinvest-cli version
```

### Go로 설치

[Go](https://go.dev/doc/install) 1.22 이상이 필요합니다.

```sh
go install github.com/isul/tossinvest-cli/cmd/tossinvest-cli@latest
```

`go install` 실행 후 바이너리는 Go bin 디렉터리에 설치됩니다.

- **기본 위치**: `$HOME/go/bin` (또는 `GOPATH/bin`)
- **경로 확인**: `go env GOPATH`

설치 후 명령이 실행되지 않으면 Go bin 디렉터리를 PATH에 추가하세요.

```sh
export PATH="$PATH:$(go env GOPATH)/bin"
```

### 로컬에서 실행

```sh
git clone https://github.com/isul/tossinvest-cli.git
cd tossinvest-cli
./scripts/run --help
```

## 사용법

리소스 기반 명령 구조:

```sh
tossinvest-cli [resource] <command> [flags...]
```

```sh
tossinvest-cli config set
tossinvest-cli prices list --symbols 005930,AAPL
tossinvest-cli holdings list
tossinvest-cli orders create --symbol 005930 --side BUY --order-type LIMIT --quantity 10 --price 70000
```

각 명령의 자세한 옵션은 `--help`로 확인하세요. 예제는 [`examples/`](examples/)에 있습니다.

## 환경 변수

| 환경 변수 | 설명 | 필수 |
| --- | --- | --- |
| `TOSSINVEST_CLIENT_ID` | OAuth client ID | 아니오 |
| `TOSSINVEST_CLIENT_SECRET` | OAuth client secret | 아니오 |
| `TOSSINVEST_ACCOUNT_SEQ` | 기본 계좌 accountSeq | 아니오 |
| `TOSSINVEST_BASE_URL` | API base URL | 아니오 (기본: `https://openapi.tossinvest.com`) |
| `TOSSINVEST_AUTO_CONFIRM` | `1`이면 주문 CONFIRM 생략 | 아니오 |

## 전역 플래그

- `--client-id` — OAuth client ID
- `--client-secret` — OAuth client secret
- `--account-seq` — `X-Tossinvest-Account` 헤더 값
- `--base-url` — 커스텀 API URL
- `--format` — 출력 형식 (`auto`, `json`, `yaml`, `pretty`, `raw`)
- `--format-error` — 오류 출력 형식
- `--transform` — [GJSON](https://github.com/tidwall/gjson) 출력 변환
- `--debug` — HTTP 디버그 로깅
- `--yes` — 쓰기 작업 CONFIRM 생략

## 주문 안전장치

`orders create`, `orders modify`, `orders cancel` 실행 전 `CONFIRM` 입력이 필요합니다.

스크립트 자동화 시 `--yes` 또는 `TOSSINVEST_AUTO_CONFIRM=1` 사용 가능 (위험 인지 필요).

## AI Agent Skill

[Agent Skills](https://agentskills.io/) 오픈 표준(`SKILL.md`) 스킬: [`skills/tossinvest/tossinvest/SKILL.md`](skills/tossinvest/tossinvest/SKILL.md). Claude Code, Codex, Copilot, Gemini CLI, OpenClaw, Hermes 등 호환 AI 코딩 에이전트에서 사용할 수 있습니다.

스킬 디렉터리를 에이전트의 skills 폴더에 복사하세요:

| 에이전트 | skills 디렉터리 |
| --- | --- |
| Claude Code | `~/.claude/skills/` 또는 `.claude/skills/` |
| OpenAI Codex | `~/.agents/skills/` 또는 `.agents/skills/` |
| GitHub Copilot | `.github/skills/` |
| Cursor | `~/.cursor/skills/` |

```sh
cp -r skills/tossinvest/tossinvest ~/.claude/skills/tossinvest
```

사용 중인 에이전트에 맞게 대상 경로를 바꿔 주세요.

## API 커버리지

[OpenAPI v1.1.1](https://openapi.tossinvest.com/openapi-docs/latest/openapi.json) 전체:

- 인증 (OAuth2, 내부 처리)
- 시세: 호가, 현재가, 체결, 상하한가, 캔들
- 종목: 기본정보, 매수 유의사항
- 시장: 환율, 국내/해외 장 캘린더
- 계좌/자산: 계좌 목록, 보유 주식
- 주문: 생성, 정정, 취소, 목록, 상세
- 주문 정보: 매수 가능 금액, 판매 가능 수량, 수수료

## Windows 11

```powershell
go install github.com/isul/tossinvest-cli/cmd/tossinvest-cli@latest
$env:Path += ";$(go env GOPATH)\bin"
tossinvest-cli version
```

## 라이선스

MIT — [LICENSE](LICENSE) 참고.
