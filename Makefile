BUILD_FOLDER=dist
BINARY_NAME=game-of-liberty

build:
	go build -o ${BUILD_FOLDER}/${BINARY_NAME}

run: build
	./${BUILD_FOLDER}/${BINARY_NAME}

# .PHONY: clean
clean:
	rm -rf ${BUILD_FOLDER}

