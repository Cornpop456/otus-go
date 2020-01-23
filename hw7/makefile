GOCMD=go
GOBUILD=$(GOCMD) build
TEST=./test.sh
BINARY_NAME=go-envdir
		
all: test build
build: 
				$(GOBUILD) -o $(BINARY_NAME) -v
test: 
				$(TEST)