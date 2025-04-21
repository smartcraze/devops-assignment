FROM golang:1.22.2-alpine


WORKDIR /app

COPY . .

RUN go build -o server .

EXPOSE 8080

CMD ["./server"]
