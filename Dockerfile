FROM golang:1.17 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /urlRedis .


FROM scratch

WORKDIR /

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /urlRedis /urlRedis

ENTRYPOINT ["/urlRedis"]
EXPOSE 8080