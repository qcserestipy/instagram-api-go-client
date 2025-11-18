SHELL := /usr/bin/env bash

.DEFAULT_GOAL := help

DOCKERCMD=$(shell which docker)
SWAGGER_VERSION=v0.30.5
SWAGGER := $(DOCKERCMD) run --rm -t -u "$(shell id -u):$(shell id -g)" -v $(shell pwd):/src -w /src quay.io/goswagger/swagger:$(SWAGGER_VERSION)

API_SPEC=api/v24.0/media/swagger.yaml
API_SPEC_ACCOUNT=api/v24.0/account/swagger.yaml
API_SPEC_PAGE=api/v24.0/page/swagger.yaml
CLIENT_DIR=pkg/sdk/v24.0/media
CLIENT_DIR_ACCOUNT=pkg/sdk/v24.0/account
CLIENT_DIR_PAGE=pkg/sdk/v24.0/page
## --------------------------------------
## Help
## --------------------------------------

help: ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z0-9_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

## --------------------------------------
## Client
## --------------------------------------

.PHONY: gen-api-client
gen-api-client: ## Generate goswagger client for media insights
	mkdir -p $(CLIENT_DIR)
	@$(SWAGGER) generate client -f ${API_SPEC} --target=$(CLIENT_DIR) --template=stratoscale --additional-initialism=CVE --additional-initialism=GC --additional-initialism=OIDC

.PHONY: gen-api-client-account
gen-api-client-account: ## Generate goswagger client for account insights
	mkdir -p $(CLIENT_DIR_ACCOUNT)
	@$(SWAGGER) generate client -f ${API_SPEC_ACCOUNT} --target=$(CLIENT_DIR_ACCOUNT) --template=stratoscale --additional-initialism=CVE --additional-initialism=GC --additional-initialism=OIDC

.PHONY: gen-api-client-page
gen-api-client-page: ## Generate goswagger client for page insights
	mkdir -p $(CLIENT_DIR_PAGE)
	@$(SWAGGER) generate client -f ${API_SPEC_PAGE} --target=$(CLIENT_DIR_PAGE) --template=stratoscale --additional-initialism=CVE --additional-initialism=GC --additional-initialism=OIDC

.PHONY: gen-all-clients
gen-all-clients: gen-api-client gen-api-client-account gen-api-client-page ## Generate all API clients

.PHONY: build
build: ## Build the main application
	go build -o bin/instagram-media-insights-go-client cmd/main/main.go

.PHONY: cleanup
cleanup: ## Clean up generated client code
	rm -rf pkg/sdk/v24.0/*
.PHONY: test
test: ## run the test
	go test ./...

.PHONY: update-submodule
update-submodule:
	@echo "Updating api submodule..."
	@cd api && git stash && git pull --rebase && cd ..
	@echo "Submodule updated successfully"

all: cleanup gen-all-clients