FROM golang:1.21 AS build

WORKDIR /app
COPY . .

RUN go mod init notion-webhook && go mod tidy
RUN go build -o server

FROM gcr.io/distroless/base-debian11
COPY --from=build /app/server /server
CMD ["/server"]
