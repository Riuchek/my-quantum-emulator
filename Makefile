.PHONY: run build test tidy clean

BIN := bin/emulator

run:
	go run ./cmd/emulator

build:
	mkdir -p bin
	go build -o $(BIN) ./cmd/emulator

test:
	go test ./...

tidy:
	go mod tidy

clean:
	rm -rf bin
