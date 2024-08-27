# Define the name of your executable and the output directory
BINARY_NAME = goat
BIN_DIR = bin

# Define the Go build flags
GOFLAGS ?= -v

# Define the Go toolchain
GO ?= go

# Ensure the bin directory exists
$(BIN_DIR):
	mkdir -p $(BIN_DIR)

# Default target
all: build

# Build the binary and place it in the bin/ directory
build: $(BIN_DIR)
	$(GO) build $(GOFLAGS) -o $(BIN_DIR)/$(BINARY_NAME) ./cmd

# Run the binary
run: build
	./$(BIN_DIR)/$(BINARY_NAME)

# Run tests
test:
	$(GO) test -v ./...

# Clean up build artifacts
clean:
	rm -f $(BIN_DIR)/$(BINARY_NAME)

# Install the binary to $GOPATH/bin
install: build
	$(GO) install

# Format Go source code
fmt:
	$(GO) fmt ./...

# Lint Go source code (requires golangci-lint to be installed)
lint:
	golangci-lint run

# Help target to display available commands
help:
	@echo "Makefile commands:"
	@echo "  all        - Build the binary"
	@echo "  build      - Build the binary"
	@echo "  run        - Build and run the binary"
	@echo "  test       - Run tests"
	@echo "  clean      - Remove build artifacts"
	@echo "  install    - Install the binary to $GOPATH/bin"
	@echo "  fmt        - Format source code"
	@echo "  lint       - Lint source code"
	@echo "  help       - Display this help message"
