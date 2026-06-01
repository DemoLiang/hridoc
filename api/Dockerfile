FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o hridoc-api hridoc.go

FROM alpine:3.19
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/hridoc-api .
COPY --from=builder /app/etc ./etc

EXPOSE 8888
CMD ["./hridoc-api", "-f", "etc/api.yaml"]
