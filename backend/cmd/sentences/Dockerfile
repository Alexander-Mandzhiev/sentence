FROM golang:alpine AS builder
WORKDIR /sentences
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd/sentences/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /sentences .
RUN chmod +x ./main
ENV CONFIG_PATH=config/sentence/production.yaml
CMD ["./main"]