GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
BINARY_NAME=calendar-app
CONFIG=./config/conf.yaml
		
all: test build
build: 
				$(GOBUILD) -o builds/$(BINARY_NAME) -v ./cmd/app
test: 
				$(GOTEST) -v ./internal/pkg/memstorage
start: 
				./builds/$(BINARY_NAME) --config $(CONFIG)
gen:
				protoc --go_out=plugins=grpc:. api/calendar/calendar.proto