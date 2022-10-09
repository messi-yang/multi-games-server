FROM golang:1.18

RUN apt update
RUN apt install -y redis-tools

COPY . /app

WORKDIR /app

RUN make build
