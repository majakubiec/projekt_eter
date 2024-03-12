FROM golang:1.21

WORKDIR /app

COPY . .

RUN go mod download && \
    go build -o main main.go

CMD ["/app/main"]