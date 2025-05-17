BINARY_NAME=bin/server
MAIN_PACKAGE=./cmd/server
MAIN_FILE=cmd/server/main.go

all: build

build:
	go build -o $(BINARY_NAME) $(MAIN_PACKAGE)

run-build:
	./$(BINARY_NAME)

run:
	go run $(MAIN_FILE)

clean:
	rm -f $(BINARY_NAME)
