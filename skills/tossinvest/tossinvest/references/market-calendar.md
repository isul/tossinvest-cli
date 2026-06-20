# Market Calendar

## get-kr

Get KR market calendar (KRX+NXT integrated mode, 3 business days).

```bash
tossinvest-cli market-calendar get-kr
```

Returns pre-market, regular, after-market session times in KST.

## get-us

Get US market calendar (3 business days).

```bash
tossinvest-cli market-calendar get-us
```

Returns `dayMarket`, `preMarket`, `regularMarket`, `afterMarket` sessions in KST. All null on holidays.

Use before placing time-sensitive orders (e.g. US amount-based market orders require regular hours).
