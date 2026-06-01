FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o hridoc-api hridoc.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/hridoc-api .
COPY --from=builder /app/etc ./etc

EXPOSE 8888
CMD ["./hridoc-api", "-f", "etc/api.yaml"]
