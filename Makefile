# Go parameters
GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOGET = $(GOCMD) get

# Name of the binary
BINARY_NAME = genji-web

# Output directory
OUTPUT_DIR = dist

# Source files
SRC_FILES = $(wildcard *.go)

# Main target: build the binary
build: $(SRC_FILES)
	@echo "Building $(BINARY_NAME)"
	$(GOBUILD) -o $(OUTPUT_DIR)/$(BINARY_NAME) .
	cp -r static $(OUTPUT_DIR)

# Run the web app
run: build
	cd $(OUTPUT_DIR) && ./$(BINARY_NAME)

# Clean up build artifacts
clean:
	$(GOCLEAN)
	rm -rf $(OUTPUT_DIR)

# Install dependencies (if needed)
deps:
	$(GOGET) ./...

# Run tests
test:
	$(GOTEST) ./...

# Build the binary and create a distribution folder
dist: clean build
	mkdir -p $(OUTPUT_DIR)
	cp -r static $(OUTPUT_DIR)

.PHONY: build run clean deps test dist
