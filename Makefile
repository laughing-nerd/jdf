FILES := $(wildcard sample/*)
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")

build-all: clean
	@mkdir -p ./dist
	@echo "Building for macOS (Intel)..."
	@GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w -X main.version=$(VERSION)" -o ./dist/jdf-darwin-amd64 .
	@echo "Building for macOS (Apple Silicon)..."
	@GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w -X main.version=$(VERSION)" -o ./dist/jdf-darwin-arm64 .
	@echo "Building for Linux (amd64)..."
	@GOOS=linux GOARCH=amd64 go build -ldflags="-s -w -X main.version=$(VERSION)" -o ./dist/jdf-linux-amd64 .
	@echo "Building for Linux (arm64)..."
	@GOOS=linux GOARCH=arm64 go build -ldflags="-s -w -X main.version=$(VERSION)" -o ./dist/jdf-linux-arm64 .
	@echo "Building for Windows (amd64)..."
	@GOOS=windows GOARCH=amd64 go build -ldflags="-s -w -X main.version=$(VERSION)" -o ./dist/jdf-windows-amd64.exe .
	@echo "Done! Binaries in ./dist/"

clean:
	@rm -rf ./dist

.PHONY: build-all clean
