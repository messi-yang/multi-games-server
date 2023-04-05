# .PHONY: postgres-plan-migrate
postgres-plan-migrate:
	migrate create -ext sql -dir db/postgres/migrations ${FILE_NAME}

# .PHONY: postgres-migrate-up
postgres-migrate-up:
	migrate -source="file:db/postgres/migrations" -database="postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" up

# .PHONY: postgres-migrate-down
postgres-migrate-down:
	migrate -source="file:db/postgres/migrations" -database="postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" down 1

# .PHONY: postgres-migrate-force
postgres-migrate-force:
	migrate -source="file:db/postgres/migrations" -database="postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" force ${POSTGRES_MIGRATE_VERSION}

# .PHONY: db-seed
db-seed:
	go run pkg/main.go db-seed

# .PHONY: dev
dev:
	air -c .air.toml

# .PHONY: lint
lint:
	golangci-lint run

# .PHONY: test
test:
	go test -v ./...

# .PHONY: build
build:
	go build -o ./dist/main ./pkg/main.go

# .PHONY: start
start:
	./dist/main

# .PHONY: clean
clean:
	rm -rf ./dist
