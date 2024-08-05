FROM golang:1.22

WORKDIR /app

COPY . .
RUN go build -o server ./src/cmd/api/main.go
CMD ["./server"]

