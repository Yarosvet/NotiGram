FROM golang:1.24.4-alpine AS build
WORKDIR /app

# скачать зависимости (кешируется при неизменных go.mod/go.sum)
COPY go.mod go.sum ./
RUN apk add --no-cache git ca-certificates && go mod download

COPY . .

# Собрать статический оптимизированный бинарник
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o /app/main .

FROM alpine:3.18 AS runtime
# установить только ca-certificates и создать непользовательского юзера
RUN apk add --no-cache ca-certificates && addgroup -S app && adduser -S -G app app

WORKDIR /home/app
COPY --from=build /app/main ./

USER app
ENTRYPOINT ["./main"]