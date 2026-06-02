# market-crawl

A market data crawling service that provides APIs for querying fund/product net value data.

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

### 2. Get CITIC Product NAV

Fetches product NAV (Net Asset Value) data from CITIC within a date range.

- **URL**: `GET /api/citic-product-nav`
- **Query Parameters**:

| Parameter   | Type   | Required | Default | Description              |
|-------------|--------|----------|---------|--------------------------|
| `prodCode`  | string | Yes      | -       | Product code             |
| `startDate` | string | Yes      | -       | Start date (YYYY-MM-DD)  |
| `endDate`   | string | Yes      | -       | End date (YYYY-MM-DD)    |

- **Example**:

```
GET /api/citic-product-nav?prodCode=xxx&startDate=2026-01-01&endDate=2026-06-01
```

---

### 3. Get BOC Revenue List

Fetches ten-thousand revenue list data from Bank of China (BOC).

- **URL**: `GET /api/boc-revenue-list`
- **Upstream**: `POST https://ebsnew.boc.cn/SAP/bocop/unlogin/ezdb/app/ten_thou_tevenue_list_info`
- **Query Parameters**:

| Parameter    | Type   | Required | Default | Description          |
|--------------|--------|----------|---------|----------------------|
| `strBakCode` | string | Yes      | -       | Product code         |
| `fundCycle`  | string | Yes      | -       | Fund cycle (e.g. 5y) |

- **Example**:

```
GET /api/boc-revenue-list?strBakCode=000509&fundCycle=5y
```
