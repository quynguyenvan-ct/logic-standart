FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o logic-app main.go

FROM alpine:latest


WORKDIR /root/

COPY --from=builder /app/logic-app .

CMD ["./logic-app", "service"]
