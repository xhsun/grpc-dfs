# Set the shell
SHELL := /bin/bash

# Set an output prefix, which is the local directory if not specified
PREFIX?=$(shell pwd)

GO_LDFLAGS_STATIC=-ldflags "-extldflags -static"

# Set default go compiler
GO := go


all: clean build-server build-cli test ## Runs a clean, build, and test

build-server: ## Builds a static executable
	@echo "+ $@"
	$(GO) build -tags "static_build" ${GO_LDFLAGS_STATIC} -o fileTransferServer ./server/cmd

build-cli: ## Builds a static executable
	@echo "+ $@"
	$(GO) build -tags "static_build" ${GO_LDFLAGS_STATIC} -o file-transfer ./client/cmd

test: ## Runs go test with coverage
	$(GO) test -coverprofile=coverage.txt  $(shell $(GO) list ./...)

docker-server: # Builds the docker images
	DOCKER_BUILDKIT="1" docker build --ssh default -f ./Dockerfile.server -t filetransferserver:latest .

clean: ## Cleanup any build binaries, packages or artifacts
	@echo "+ $@"
	$(RM) fileTransferServer
	$(RM) file-transfer
	$(RM) coverage.txt