FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod .
RUN go mod download

COPY . /app

ENV SERVICE=tw_TIMELINE
ENV ID=tw_TIMELINE_1
ENV PORT=8100
ENV EXCHNAME=Timeline_tweet
ENV QNAME=Timeline
ENV CONSUMER=Timeline_CONSUMER1

RUN go build -o main ./cmd/main.go

EXPOSE 8100

CMD ["./main"]