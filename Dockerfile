#Build
FROM golang:1.24.4-alpine AS Builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /ai-api ./cmd/ai-agent/main.go

#Run
FROM alpine:3.20
RUN adduser -D appuser
COPY --from=builder /ai-api /ai-api
USER appuser
EXPOSE 8080
ENTRYPOINT ["/ai-api"]
