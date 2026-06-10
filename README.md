# market-crawl

A market data crawling service that provides APIs for querying fund/product net value data.

## Disclaimer

This repository is an independent personal project and has no official, authorized, or affiliated relationship with the companies referenced by prior upstream service examples or related entities.

## Running

```bash
go run main.go
```

The server starts on port `8080`.

## Docker

```bash
docker build -t market-crawl .
docker run -p 8080:8080 market-crawl
```

## APIs

### 1. Get Net Value List

Fetches net value data for a given product (ICBC source).

- **URL**: `GET /api/net-value-list`
- **Query Parameters**:

| Parameter   | Type   | Required | Default | Description          |
|-------------|--------|----------|---------|----------------------|
| `prodId`    | string | Yes      | -       | Product ID           |
| `pageIndex` | int    | No       | 1       | Page number          |
| `pageSize`  | int    | No       | 10      | Number of items per page |

- **Example**:

```
GET /api/net-value-list?prodId=xxx&pageIndex=1&pageSize=10
```

---

### 2. Get BOC Revenue List

Fetches ten-thousand revenue list data from Bank of China (BOC) and calculates NAV for each item based on a known base date and NAV value.

- **URL**: `GET /api/boc-revenue-list`
- **Upstream**: `POST https://ebsnew.boc.cn/SAP/bocop/unlogin/ezdb/app/ten_thou_tevenue_list_info`
- **Query Parameters**:

| Parameter    | Type   | Required | Default | Description                                      |
|--------------|--------|----------|---------|--------------------------------------------------|
| `strBakCode` | string | Yes      | -       | Product code                                     |
| `fundCycle`  | string | Yes      | -       | Fund cycle (e.g. 5y)                             |
| `baseDate`   | string | Yes      | -       | Date with known NAV (YYYY-MM-DD)                 |
| `baseNav`    | float  | Yes      | -       | Known NAV on the base date (e.g. 1.0344)         |

- **NAV Calculation** (dividend-reinvestment model, 红利再投资):
  - `tenThouRet` is the absolute daily revenue per 10,000 shares (CNY); under
    reinvestment the per-share daily increment is `tenThouRet / 10000`.
  - Forward:  `nav[i] = nav[i-1] + tenThouRet[i]   / 10000`
  - Backward: `nav[i] = nav[i+1] - tenThouRet[i+1] / 10000`

- **Example**:

```
GET /api/boc-revenue-list?strBakCode=000509&fundCycle=5y&baseDate=2026-06-01&baseNav=1.0344
```
