FROM golang:1.24-alpine AS builder

WORKDIR /app

RUN apk add --no-cache gcc musl-dev

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o jamlink ./cmd/api

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/jamlink .

EXPOSE 8080

CMD ["./jamlink"]
