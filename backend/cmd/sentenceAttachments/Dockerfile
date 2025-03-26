FROM golang:alpine AS builder
WORKDIR /sentence_attachments
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd/sentenceAttachments/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /sentence_attachments .
RUN chmod +x ./main
ENV CONFIG_PATH=config/sentence_attachments/production.yaml
CMD ["./main"]