# Этап, на котором выполняется сборка приложения
FROM golang:1.19-alpine as builder
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
#RUN go mod tidy
RUN go build -o bin/app-serverHttp app/main.go
# Финальный этап, копируем собранное приложение
FROM alpine:latest
COPY --from=builder /build/bin/app-serverHttp /bin/app-serverHttp
ENV NEWS_DB_NAME=postgres
ENV NEWS_DB_PASSWORD=postgres
ENV NEWS_DB_USER=postgres
ENV NEWS_DB_HOST=172.17.0.2
ENV NEWS_DB_PORT=5432
ENV NEWS_HTTP_PORT=8001
ENV NEWS_HTTP_HOST_LISTEN=0.0.0.0
EXPOSE $NEWS_HTTP_PORT
ENTRYPOINT ["/bin/app-serverHttp"]