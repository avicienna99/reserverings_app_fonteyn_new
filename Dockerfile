FROM golang:1.23-alpine

WORKDIR /app

COPY . .

RUN go mod download && go build -o server

EXPOSE 80

CMD ["./server"]