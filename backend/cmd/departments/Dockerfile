FROM golang:alpine AS builder
WORKDIR /departments
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd/departments/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /departments .
RUN chmod +x ./main
ENV CONFIG_PATH=config/departments/production.yaml
CMD ["./main"]