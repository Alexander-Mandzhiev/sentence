FROM golang:alpine AS builder
WORKDIR /direction
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd/direction/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /direction .
RUN chmod +x ./main
ENV CONFIG_PATH=config/directions/production.yaml
CMD ["./main"]