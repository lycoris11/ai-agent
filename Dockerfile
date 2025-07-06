#Build
FROM golang:1.24.4-alpine AS Builder

WORKDIR /app

RUN apk add --no-cache git
RUN git clone https://github.com/lycoris11/ai-agent

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /ai-api ./cmd/ai-agent/main.go

#Run
FROM alpine:3.20
RUN adduser -D appuser
COPY --from=builder /ai-api /ai-api
USER appuser
EXPOSE 8080
ENTRYPOINT ["/ai-api"]
