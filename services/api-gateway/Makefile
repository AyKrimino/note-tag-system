.PHONY: build run test clean

BINARY=api-gateway
SRC=main.go

build:
	@go build -o bin/$(BINARY) $(SRC)

run: build
	@./bin/$(BINARY)

test:
	@go test ./...

clean:
	@rm -rf bin
