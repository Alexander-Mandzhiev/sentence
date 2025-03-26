FROM golang:alpine AS builder
WORKDIR /history
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd/history/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /history .
RUN chmod +x ./main
ENV CONFIG_PATH=config/history/production.yaml
CMD ["./main"]