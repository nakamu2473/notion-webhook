FROM golang:1.24 AS build

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o server && chmod +x server

FROM gcr.io/distroless/static-debian11
COPY --from=build /app/server /server
CMD ["/server"]
