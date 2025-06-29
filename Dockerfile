FROM golang:1.24 AS build

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

RUN go build -o server && chmod +x server

#FROM gcr.io/distroless/base-debian11
FROM debian:bullseye-slim

COPY --from=build /app/server /server

CMD ["/server"]