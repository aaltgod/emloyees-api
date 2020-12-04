FROM golang:1.15

MAINTAINER tg: @alaskastorm

RUN mkdir /app

ADD . /app

WORKDIR /app

#RUN go mod download

RUN go build -o rest-api

CMD ["/app/rest-api"]