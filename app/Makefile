BINARY_NAME=lol-team-maker
BUILD_DIR=bin

all: build

build:
	@echo "Building the binary..."
	go build -o $(BUILD_DIR)/$(BINARY_NAME) .

clean:
	@echo "Cleaning up..."
	rm -f $(BUILD_DIR)/$(BINARY_NAME)