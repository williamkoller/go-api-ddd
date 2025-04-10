FROM golang:1.24.2

WORKDIR /app
COPY . .

RUN go build -o main ./cmd/main.go

RUN apt-get update && apt-get install -y netcat-openbsd

RUN chmod +x wait-for-it.sh

CMD ["./wait-for-it.sh", "db", "3306", "--", "./main"]
