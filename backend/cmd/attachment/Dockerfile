FROM golang:alpine AS builder
WORKDIR /attachment
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd/attachment/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /attachment .
RUN chmod +x ./main
ENV CONFIG_PATH=config/attachments/production.yaml
CMD ["./main"]