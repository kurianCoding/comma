BIN=$(HOME)/bin
test:
	comma  hello.sh
build:
	go build -o $(BIN)/comma
