BUILD_FOLDER=dist
BINARY_NAME=game-of-liberty

# .PHONY: init-cassandra
init-cassandra:
	cqlsh ${CASSANDRA_HOST} -f ./db/cassandra/init.cql

# .PHONY: connect-cassandra
connect-cassandra:
	cqlsh ${CASSANDRA_HOST}

# .PHONY: plan-cassandra-migrate
plan-cassandra-migrate:
	migrate create -ext sql -dir db/cassandra/migrations ${FILE_NAME}

# .PHONY: cassandra-migrate-up
cassandra-migrate-up:
	migrate -source="file:db/cassandra/migrations" -database="cassandra://${CASSANDRA_HOST}:${CASSANDRA_PORT}/${CASSANDRA_KEYSPACE}?x-multi-statement=true" up

# .PHONY: cassandra-migrate-down
cassandra-migrate-down:
	migrate -source="file:db/cassandra/migrations" -database="cassandra://${CASSANDRA_HOST}:${CASSANDRA_PORT}/${CASSANDRA_KEYSPACE}?x-multi-statement=true" down

# .PHONY: plan-postgres-migrate
plan-postgres-migrate:
	migrate create -ext sql -dir db/postgres/migrations ${FILE_NAME}

# .PHONY: postgres-migrate-up
postgres-migrate-up:
	migrate -source="file:db/postgres/migrations" -database="postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" up

# .PHONY: postgres-migrate-down
postgres-migrate-down:
	migrate -source="file:db/postgres/migrations" -database="postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" down

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
