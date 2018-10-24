PKG_NAME := go-rest-api
PKG_PATH := /go/src/$(PKG_NAME)

.PHONY: help all start stop build run reset ssh

help: ## Display helpful info
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

all: start build run

start: ## Start the app container
	docker run --rm -id --name $(PKG_NAME) -w $(PKG_PATH) -v $$(pwd):$(PKG_PATH) -p 127.0.0.1:8888:80 golang:1.11.1 /bin/bash

stop: ## Stop the app container
	docker stop $(PKG_NAME)

build: ## Build the app binary
	docker exec -it $(PKG_NAME) go build

run: ## Run the app
	docker exec -it $(PKG_NAME) $(PKG_PATH)/$(PKG_NAME)

reset: build run ## Rebuild and run the app

ssh: ## Enter the app container
	docker exec -it $(PKG_NAME) /bin/bash
