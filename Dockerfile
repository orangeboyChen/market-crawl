# Build stage
FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/market-crawl .

# Runtime stage
FROM gcr.io/distroless/static-debian12:nonroot

COPY --from=builder /app/market-crawl /market-crawl

EXPOSE 8080

ENTRYPOINT ["/market-crawl"]
