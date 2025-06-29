FROM golang:1.24 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main .

# 本番用の軽量イメージにバイナリだけコピーする
FROM gcr.io/distroless/base-debian11

WORKDIR /
COPY --from=builder /app/main /

# Cloud Run は $PORT を自動で設定する
ENV PORT=8080
EXPOSE 8080

CMD ["/main"]
