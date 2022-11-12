BUILD_FOLDER=dist
BINARY_NAME=game-of-liberty

# .PHONY: create-db-migrate-file
create-db-migrate:
	migrate create -ext sql -dir db/migration $(FILE_NAME)

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
