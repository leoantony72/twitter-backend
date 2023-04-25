FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod .
RUN go mod download

COPY . /app

ENV SERVICE=tw_Auth
ENV ID=tw_Auth_1
ENV PORT=8080

RUN go run main.go
RUN go build -o main ./cmd/main.go

EXPOSE 8080

CMD ["./main"]

