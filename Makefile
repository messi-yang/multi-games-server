BUILD_FOLDER=dist
BINARY_NAME=game-of-liberty

# .PHONY: create-db-migrate
create-migrate-db-file:
	migrate create -ext sql -dir db/migration $(FILE_NAME)

# .PHONY: db-migrate
migrate-db:
	migrate -source file:db/migration -database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:5432/${POSTGRES_DB}?sslmode=disable up

# .PHONY: build
build:
	go build -o ${BUILD_FOLDER}/${BINARY_NAME}

# .PHONY: dev
dev:
	air

# .PHONY: start
start:
	./${BUILD_FOLDER}/${BINARY_NAME}

# .PHONY: clean
clean:
	rm -rf ${BUILD_FOLDER}
