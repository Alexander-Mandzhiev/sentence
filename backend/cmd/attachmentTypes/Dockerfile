FROM golang:alpine AS builder
WORKDIR /attachmentTypes
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd/attachmentTypes/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /attachmentTypes .
RUN chmod +x ./main
ENV CONFIG_PATH=config/attachment_types/production.yaml
CMD ["./main"]