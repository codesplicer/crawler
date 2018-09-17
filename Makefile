# Self-documenting Makefile: https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html

PACKAGE = github.com/codesplicer/crawler
GOEXE ?= go

.PHONY: help
.DEFAULT_GOAL := help

vendor: ## Install dep and sync vendored dependencies
		${GOEXE} get github.com/golang/dep
		dep ensure ${PACKAGE}

test: ## Run tests
		${GOEXE} test -v

build: vendor ## Build crawler binary
		${GOEXE} build -o bin/crawler

help:
		@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
