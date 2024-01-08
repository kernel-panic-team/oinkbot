FROM golang:alpine AS builder
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . .
ENV CGO_ENABLED=0 GO111MODULE=on
RUN go build -o ./bin ./cmd/main.go

FROM alpine:latest
#VOLUME /ssl
WORKDIR /
COPY --from=builder /app/bin/* ./
RUN #chmod +x ./main
ENTRYPOINT ["./main"]