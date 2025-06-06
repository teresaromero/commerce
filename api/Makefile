.PHONY: run build test vet

# Default binary output path
BINARY=bin/commerce-api

# Run the application
run:
	go run cmd/main.go

# Build the binary
build:
	go build -o $(BINARY) cmd/main.go

# Run unit tests
test:
	go test ./... -v

# Run go vet for static analysis
vet:
	go vet ./...

# Clean binary
clean:
	rm -f $(BINARY)