# Glossary (한영 용어)

## Trading

| English | Korean |
|---|---|
| Buy | 매수 |
| Sell | 매도 |
| Order | 주문 |
| Limit order | 지정가 주문 |
| Market order | 시장가 주문 |
| Fill / Execution | 체결 |
| Open order | 미체결 주문 |
| Cancel | 취소 |
| Modify / Replace | 정정 |
| Holdings | 보유 주식 |
| Buying power | 매수 가능 금액 |
| Sellable quantity | 판매 가능 수량 |
| Commission | 수수료 |
| Orderbook | 호가 |
| Candle | 캔들 (봉) |
| Exchange rate | 환율 |

## API Fields

| Field | Korean | Description |
|---|---|---|
| `symbol` | 종목 심볼 | KR: 6-digit (`005930`), US: ticker (`AAPL`) |
| `side` | 주문 방향 | `BUY` / `SELL` |
| `orderType` | 호가 유형 | `LIMIT` / `MARKET` |
| `timeInForce` | 유효 조건 | `DAY` / `CLS` |
| `quantity` | 수량 | Integer shares |
| `orderAmount` | 주문 금액 | USD amount (US market only) |
| `price` | 가격 | Limit price |
| `accountSeq` | 계좌 일련번호 | Used in `X-Tossinvest-Account` header |
| `orderId` | 주문 ID | Server-issued opaque token |
| `status` | 주문 상태 | See orders reference |

## Order Status

| Status | Korean |
|---|---|
| `PENDING` | 대기 |
| `PARTIAL_FILLED` | 부분 체결 |
| `FILLED` | 전량 체결 |
| `CANCELED` | 취소됨 |
| `REJECTED` | 거부됨 |
| `PENDING_CANCEL` | 취소 대기 |
| `PENDING_REPLACE` | 정정 대기 |
| `REPLACED` | 정정 완료 |

## List Filter Groups

| Filter | Includes |
|---|---|
| `OPEN` | In-progress orders |
| `CLOSED` | Completed/terminated orders |
