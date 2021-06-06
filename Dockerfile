FROM golang:1.12.0-alpine3.9

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go build -o main.go

CMD ["./api-amalan-nahdliyin"]