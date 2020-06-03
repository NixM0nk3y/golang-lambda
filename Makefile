# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=golang-lambda

VERSION=1.0.0

# source database
SAMBUCKET ?= aws-sam-cli-managed-default-samclisourcebucket-1bac5dz9xjbuu

COMMIT=$(shell git rev-list -1 HEAD --abbrev-commit)
DATE=$(shell date -u '+%Y%m%d')

all: test lambda/build

deps:
	go get -v  ./...

lambda/build:
	$(GOBUILD) -ldflags " \
		-X github.com/NixM0nk3y/golang-lambda/pkg/version.Version=${VERSION} \
		-X github.com/NixM0nk3y/golang-lambda/pkg/version.BuildHash=${COMMIT} \
		-X github.com/NixM0nk3y/golang-lambda/pkg/version.BuildDate=${DATE}" \
		-o ./bin/${BINARY_NAME} -v ./cmd/${BINARY_NAME}

lambda/test: lambda/build
	sam local invoke "golangFunction" --env-vars ./test/testenvironment.json --event ./test/request.json

lambda/panic: lambda/build
	sam local invoke "golangFunction" --env-vars ./test/testenvironment.json --event ./test/panicrequest.json

test/lamda/start:
	sam local start-lambda --env-vars ./test/testenvironment.json

# no worky - https://github.com/awslabs/aws-sam-cli/issues/1641
test/api/start:
	sam local start-api

test: 
		$(GOTEST) -v ./...
clean: 
		$(GOCLEAN)
		rm -f ./bin/*

