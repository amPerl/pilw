GOCMD=go
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

BINARY_NAME=pilw
BINARY_ENTRYPOINT=cmd/pilw/main.go

run:
	$(GORUN) $(BINARY_ENTRYPOINT)

clean:
	$(GOCLEAN) $(BINARY_ENTRYPOINT)
	rm $(BINARY_NAME)

build:
	$(GOBUILD) -o $(BINARY_NAME) $(BINARY_ENTRYPOINT)

install: build
	cp $(BINARY_NAME) /usr/local/bin/$(BINARY_NAME)
	chmod +x /usr/local/bin/$(BINARY_NAME)
