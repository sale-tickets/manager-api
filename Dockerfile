FROM golang:1.23-alpine AS builder

WORKDIR /build-go

COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/main.go

FROM alpine:latest

WORKDIR /run

COPY --from=builder /build-go/ /run/

EXPOSE 8000
EXPOSE 50000

CMD [ "./app" ]