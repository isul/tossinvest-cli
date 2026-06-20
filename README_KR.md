[English](./README.md) | 한국어

# tossinvest-cli

[토스증권 Open API](https://developers.tossinvest.com/docs)를 위한 CLI입니다.

> **이 프로그램은 토스증권 공식 프로그램이 아닙니다.**

## 설치

### npm으로 설치

```sh
npm install -g @isul/tossinvest-cli
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

Cursor 등 AI 도구용 스킬: [`skills/tossinvest/tossinvest/SKILL.md`](skills/tossinvest/tossinvest/SKILL.md)

```sh
mkdir -p ~/.cursor/skills
cp -r skills/tossinvest/tossinvest ~/.cursor/skills/tossinvest
```

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

```powershell
npm install -g @isul/tossinvest-cli
tossinvest-cli version
```

## 라이선스

MIT — [LICENSE](LICENSE) 참고.
