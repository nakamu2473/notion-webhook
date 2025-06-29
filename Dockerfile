FROM golang:1.24 AS build

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o server

FROM gcr.io/distroless/base-debian11
COPY --from=build /app/server /server
CMD ["/server"]
