go build: example

all: clean dependencies build

re: clean build

clean:
	@echo "Cleaning..."
	-rm -f $(BINARY_NAME)
	-rm -f $(BINARY_UNIX)
	-find . -type d -name -exec rm -rf \{} +
	-$(GOCLEAN) -i
	@echo "Done cleaning."

dependencies:
	@echo "Getting dependencies..."
	$(GOMOD) download
	$(GOMOD) vendor
	@echo "Done getting dependencies."

example:
	@echo "Building Example"
#	@echo "Installing vendors..."
#	go install ./vendor/...
	@echo "Building..."
	$(GOBUILD) ./...
	$(GOBUILD) -o $(BINARY_NAME) -v cmd/*.go

serve:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME) serve

update-dependencies: dependencies
	$(GOGET) -u

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v cmd/*.go

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOLIST=$(GOCMD) list
GOMOD=GO111MODULE=on $(GOCMD) mod
GOGET=GO111MODULE=on $(GOCMD) get
BINARY_NAME=example
BINARY_UNIX=$(BINARY_NAME)_unix
