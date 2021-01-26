# Parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
BINARY_NAME=gowars

all: build

clean:
    $(GOCLEAN)
    rm -rf bin

assets:
    go-bindata schema

build: assets
    CGO_ENABLED=0 $(GOBUILD) -o bin/$(BINARY_NAME) -v
