.PHONY: build clean install

# Binary name
BINARY=aa

# Build directory
BUILD_DIR=build

# Go build flags
LDFLAGS=-s -w

build:
	@echo "Building..."
	@mkdir -p $(BUILD_DIR)
	@go build -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/$(BINARY)

install: build
	@echo "Installing..."
	@mv $(BUILD_DIR)/$(BINARY) $(GOPATH)/bin/$(BINARY)

clean:
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)
	@rm -f $(GOPATH)/bin/$(BINARY)

test:
	@go test -v ./...
