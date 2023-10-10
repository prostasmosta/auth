FROM golang:1.21.1-alpine AS builder

COPY . /github.com/prostasmosta/auth/grpc/source/
WORKDIR /github.com/prostasmosta/auth/grpc/source/

RUN go mod download
RUN go build -o ./bin/crud_server cmd/grpc_server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/prostasmosta/auth/grpc/source/bin/crud_server .

CMD ["./crud_server"]
