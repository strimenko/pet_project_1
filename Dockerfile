# Этап сборки
FROM golang:1.24 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Сборка статического бинарника
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp ./cmd

# Финальный образ
FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/myapp /usr/local/bin/myapp
COPY .env .env


CMD ["myapp"]
