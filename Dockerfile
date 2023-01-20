FROM golang:1.18

# Install "redis"
RUN apt update
RUN apt install -y redis-tools

# Install "migrate" https://github.com/golang-migrate/migrate
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.15.2

# Install "air"
RUN go install github.com/cosmtrek/air@latest

COPY . /app

WORKDIR /app

RUN make build-api-server
RUN make build-game-server
