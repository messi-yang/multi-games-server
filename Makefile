BUILD_FOLDER=dist
BINARY_NAME=game-of-liberty

# .PHONY: kafka
kafka:
	docker-compose up

# .PHONY: install-air
install-air:
	go install github.com/cosmtrek/air@latest

# .PHONY: build
build:
	go build -o ${BUILD_FOLDER}/${BINARY_NAME}

# .PHONY: dev
dev: install-air
	air

# .PHONY: start
start: build
	./${BUILD_FOLDER}/${BINARY_NAME}

# .PHONY: clean
clean:
	rm -rf ${BUILD_FOLDER}

