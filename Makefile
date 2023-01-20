BUILD_FOLDER=dist
BINARY_NAME=game-of-liberty

# .PHONY: create-db-migrate
create-migrate-db-file:
	migrate create -ext sql -dir db/migration $(FILE_NAME)

# .PHONY: db-migrate
migrate-db:
	migrate -source file:db/migration -database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:5432/${POSTGRES_DB}?sslmode=disable up

# .PHONY: api-server-dev
api-server-dev:
	air -c .air.api-server.toml

# .PHONY: game-server-dev
game-server-dev:
	air -c .air.game-server.toml

# .PHONY: build-api-server
build-api-server:
	go build -o ./dist/api-server ./src/apiserver/main.go

# .PHONY: build
build-game-server:
	go build -o ./dist/game-server ./src/gameserver/main.go

# .PHONY: start-api-server
start-api-server:
	./dist/api-server

# .PHONY: start-api-server
start-game-server:
	./dist/game-server

# .PHONY: clean
clean-builds:
	rm -rf ./dist
