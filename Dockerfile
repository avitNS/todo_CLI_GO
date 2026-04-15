FROM golang:1.25-alpine AS builder

WORKDIR /build

COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /todo-cli ./cmd/main.go


FROM alpine:3.22

WORKDIR /app

COPY --from=builder /todo-cli .

RUN addgroup -S gouser && adduser -S gouser -G gouser

RUN mkdir -p /app/data && chown -R gouser:gouser /app/data

ENV TODO_STORAGE_PATH=/app/data/tasks.json

ENTRYPOINT [ "./todo-cli" ]
