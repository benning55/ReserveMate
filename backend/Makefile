BINARY_NAME=server
.DEFAULT_GOAL := run

build:
	go build -o ./target/${BINARY_NAME} main.go

run: build
	./target/${BINARY_NAME}

clean:
	go clean
	rm ./target/${BINARY_NAME}