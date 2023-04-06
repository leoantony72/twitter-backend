FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod .
RUN go mod download

COPY . /app

ENV SERVICE=tw_Tweet
ENV ID=tw_Tweet_1
ENV PORT=8090
ENV EXCHNAME=Timeline_tweet

RUN go build -o main ./cmd/main.go

EXPOSE 8090

CMD ["./main"]

