FROM golang:alpine AS builder
WORKDIR /gateway
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd/gateway/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /gateway .
RUN chmod +x ./main
ENV CONFIG_PATH=config/gateway/production.yaml
CMD ["./main"]