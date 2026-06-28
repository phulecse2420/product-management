FROM golang:1.26.4 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o server ./main.go


FROM alpine:3.19

WORKDIR /app

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

COPY --from=builder /app/server .
COPY --chown=appuser:appgroup migrations/ ./migrations/

USER appuser

EXPOSE 8080

CMD ["./server"]