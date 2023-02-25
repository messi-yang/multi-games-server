BUILD_FOLDER=dist
BINARY_NAME=game-of-liberty

# .PHONY: init-cassandra
init-cassandra:
	cqlsh ${CASSANDRA_HOST} -f ./db/cassandra/init.cql

# .PHONY: connect-cassandra
connect-cassandra:
	cqlsh ${CASSANDRA_HOST}

# .PHONY: create-cassandra-migrate-file
create-cassandra-migrate-file:
	migrate create -ext sql -dir db/cassandra/migrations ${FILE_NAME}

# .PHONY: start-cassandra-migrate
start-cassandra-migrate:
	migrate -source="file:db/cassandra/migrations" -database="cassandra://${CASSANDRA_HOST}:${CASSANDRA_PORT}/${CASSANDRA_KEYSPACE}?x-multi-statement=true" up

# .PHONY: revert-cassandra-migrate
revert-cassandra-migrate:
	migrate -source="file:db/cassandra/migrations" -database="cassandra://${CASSANDRA_HOST}:${CASSANDRA_PORT}/${CASSANDRA_KEYSPACE}?x-multi-statement=true" down

# .PHONY: dev
dev:
	air -c .air.toml

# .PHONY: build
build:
	go build -o ./dist/main ./pkg/main.go

# .PHONY: start
start:
	./dist/main

# .PHONY: clean-build
clean-build:
	rm -rf ./dist
