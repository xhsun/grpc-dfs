# Set the shell
SHELL := /bin/bash

# Set an output prefix, which is the local directory if not specified
PREFIX?=$(shell pwd)

GO_LDFLAGS_STATIC=-ldflags "-extldflags -static"

# Set default go compiler
GO := go


all: clean build-server build-cli test ## Runs a clean, build, and test

build: build-server build-cli

build-server: ## Builds a static executable
	@echo "+ $@"
	$(GO) build -tags "static_build" ${GO_LDFLAGS_STATIC} -o dfsServer ./server/cmd

build-cli: ## Builds a static executable
	@echo "+ $@"
	$(GO) build -tags "static_build" ${GO_LDFLAGS_STATIC} -o dfs-cli ./client/dfs-cli

test: ## Runs go test with coverage
	$(GO) test -coverprofile=coverage.txt  $(shell $(GO) list ./...)

install-cli: ## Build and install a static executable
	@echo "+ $@"
	$(GO) install -tags "static_build" ${GO_LDFLAGS_STATIC} ./client/dfs-cli

docker-server: # Builds the docker images
	DOCKER_BUILDKIT="1" docker build --ssh default -f ./Dockerfile.server -t dfs-server:latest .

clean: ## Cleanup any build binaries, packages or artifacts
	@echo "+ $@"
	$(RM) dfsServer
	$(RM) dfs-cli
	$(RM) coverage.txt