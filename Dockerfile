FROM golang:1.20

# Install "redis"
RUN apt update
RUN apt install -y redis-tools

# Install "cqlsh" that is shipped with "cassandra"
RUN apt install openjdk-11-jdk -y
RUN apt install apt-transport-https gnupg2 -y
RUN wget -q -O - https://www.apache.org/dist/cassandra/KEYS | apt-key add -
RUN sh -c 'echo "deb https://downloads.apache.org/cassandra/debian 40x main" | tee -a /etc/apt/sources.list.d/cassandra.sources.list'
RUN apt update
RUN apt install cassandra -y

# Install "migrate" https://github.com/golang-migrate/migrate
RUN go install -tags 'cassandra postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Install "air"
RUN go install github.com/cosmtrek/air@latest

# Install Golang CLI Lint
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.52.1


COPY . /app

WORKDIR /app

RUN make build
