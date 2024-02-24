FROM golang:1.22

# Install "redis"
RUN apt update
RUN apt install -y redis-tools postgresql-client

# Install "migrate" https://github.com/golang-migrate/migrate
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Install "air"
RUN go install github.com/cosmtrek/air@latest

# Install Golang CLI Lint
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.52.1

COPY . /app

WORKDIR /app

RUN make build

CMD ["sh", "-c", "./bin/start.sh"]
