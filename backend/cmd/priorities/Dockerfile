FROM golang:alpine AS builder
WORKDIR /priorities
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd/priorities/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /priorities .
RUN chmod +x ./main
ENV CONFIG_PATH=config/priorities/production.yaml
CMD ["./main"]