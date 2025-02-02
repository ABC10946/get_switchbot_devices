# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build -trimpath
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=get_switchbot_devices
BINARY_UNIX=$(BINARY_NAME)_unix

# Build the project
build:
	CGO_ENABLED=0 $(GOBUILD) -o $(BINARY_NAME) -v

# Clean the project
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

# Run tests
test:
	$(GOTEST) -v ./...

# Install dependencies
deps:
	$(GOGET) -u github.com/nasa9084/go-switchbot

# Cross compilation for Linux
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v

.PHONY: build clean test deps build-linux

install:
	cp $(BINARY_NAME) $(HOME)/.local/bin/$(BINARY_NAME)