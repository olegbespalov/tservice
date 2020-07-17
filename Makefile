.EXPORT_ALL_VARIABLES:
GO111MODULE=on
GOCMDWITHOUTMODULES = GO111MODULE="off" GOFLAGS="" $(GOCMD)
TEST_RESULTS_DIR=./test/results

GOCMD = $(shell which go)
GOFMT = $(shell which gofmt)
GOFILES = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

.PHONY: help
help: ## - Show help message
	@printf "\033[32m\xE2\x9c\x93 usage: make [target]\n\n\033[0m"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: app_build
app_build: ## - Build a TService localy on a host machine
	$(GOCMD) build -o bin/tservice cmd/tservice/main.go

.PHONY: code_format
code_format: ## - Format code
	$(GOFMT) -d -w $(GOFILES)

.PHONY: code_lint
code_lint: ## - Run a linter
	(which golangci-lint || $(GOCMDWITHOUTMODULES) get github.com/golangci/golangci-lint/cmd/golangci-lint)
	golangci-lint run	

.PHONY: code_coverage
code_coverage: ## - measure the code test coverage
	(which $(GOCMD) tool cover || $(GOCMDWITHOUTMODULES) get golang.org/x/tools/cmd/cover)
	$(GOCMD) test ./... -coverprofile=$(TEST_RESULTS_DIR)/coverage.out

.PHONY: code_test
code_test: ## - Run tests
	$(GOCMD) test ./... -short
