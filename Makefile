BUILD_FOLDER=dist
BINARY_NAME=game-of-liberty

# .PHONY: create-db-migrate
create-migrate-db-file:
	migrate create -ext sql -dir db/migration $(FILE_NAME)

# .PHONY: db-migrate
migrate-db:
	migrate -source file:db/migration -database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:5432/${POSTGRES_DB}?sslmode=disable up

# .PHONY: dev
dev:
	air -c .air.toml

# .PHONY: build
build:
	go build -o ./dist/main ./pkg/main.go

# .PHONY: start
start:
	./dist/main

# .PHONY: clean
clean-builds:
	rm -rf ./dist
