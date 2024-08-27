# Define the name of your executable and the output directory
BINARY_NAME = goat
BIN_DIR = bin

# Ensure the bin directory exists
$(BIN_DIR):
	mkdir -p $(BIN_DIR)

# Default target: build the binary
all: build

# Build the binary and move it to the bin/ directory
build: $(BIN_DIR)
	go build -o $(BINARY_NAME) .
	mv $(BINARY_NAME) $(BIN_DIR)/

# Run the binary with optional arguments
run: build
	$(BIN_DIR)/$(BINARY_NAME) $(ARGS)

# Clean up build artifacts
clean:
	rm -f $(BIN_DIR)/$(BINARY_NAME)

# Help target to display available commands
help:
	@echo "Makefile commands:"
	@echo "  all        - Build the binary"
	@echo "  build      - Build the binary"
	@echo "  run        - Build and run the binary with optional arguments"
	@echo "  clean      - Remove build artifacts"
	@echo "  help       - Display this help message"
