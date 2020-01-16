GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
BINARY_NAME=gocopy
		
all: test build
build: 
				$(GOBUILD) -o $(BINARY_NAME) -v
test: 
				$(GOTEST) -v ./...