# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
MAIN_PATH=./
BINARY_DIR=./bin
BINARY_NAME=attacker
BINARY_PATH=$(BINARY_DIR)/$(BINARY_NAME)

help:			## show targets list
	@echo Targets:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build-mac: ## build mac binary
	@GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 $(GOBUILD) -o $(BINARY_DIR)/mac64/$(BINARY_NAME) -v $(MAIN_PATH)

build-linux: ## build linux binary
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 $(GOBUILD) -o $(BINARY_DIR)/linux64/$(BINARY_NAME) -v $(MAIN_PATH)

build-all-os: build-mac build-linux ## build all os version binary
	@echo build-all

clean:			## clean up binary folder
	@if [ -d $(BINARY_DIR) ] ; then rm -r $(BINARY_DIR)/* ; fi

.PHONY: help build-mac build-linux build-all-os clean