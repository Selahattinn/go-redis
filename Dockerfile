FROM golang:1.16-alpine


WORKDIR /app

COPY go.* ./

RUN go mod download
COPY . ./

RUN go build ./cmd/go-redis/go-redis.go

EXPOSE 8080

CMD ["./go-redis"]