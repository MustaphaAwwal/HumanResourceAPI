FROM golang:1.21.0

WORKDIR /app

COPY . .

RUN go build -o api

EXPOSE 80

CMD ["./api"]
