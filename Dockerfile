FROM golang:1.18

COPY . /app

WORKDIR /app

RUN make build

CMD make start
