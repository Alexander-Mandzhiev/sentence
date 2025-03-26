FROM golang:alpine AS builder
WORKDIR /implementors
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd/implementors/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /implementors .
RUN chmod +x ./main
ENV CONFIG_PATH=config/implementors/production.yaml
CMD ["./main"]