GO=go
GOCOVER=$(GO) tool cover
GOTEST=$(GO) test
COVERFILE=coverage.out

GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)

.PHONY: test cover

all: help

## Testing:

vet: ## Run vet
	$(GO) vet ./...

test: ## Run all tests with race detection
	$(GOTEST) -v -race ./...

testfast: ## Run all tests
	$(GOTEST) -v ./...

cover: ## Run tests and open coverage in browser
	$(GOTEST) -v -coverpkg=./... -covermode=atomic -coverprofile=$(COVERFILE) ./...
	$(GOCOVER) -func=$(COVERFILE)
	$(GOCOVER) -html=$(COVERFILE)
	@rm $(COVERFILE)

coverall: ## Run tests and print coverage across all packages
	$(GOTEST) -v -coverpkg=./... -coverprofile=$(COVERFILE) ./... 
	$(GOCOVER) -func $(COVERFILE) 
	@rm $(COVERFILE)

## Build & Run:

build: ## Statically compile application to dist/
	CGO_ENABLED=0 $(GO) build -ldflags="-w -s" -v -o ./dist/nextversion ./cmd/nextversion/

run: ## Compile and run application (development)
	@CGO_ENABLED=0 $(GO) run ./cmd/nextversion/ $(ARGS)

run-race: ## Compile and run application with race detection (development)
	@CGO_ENABLED=0 $(GO) run -race ./cmd/nextversion/ $(ARGS)

## Help:
help: ## Show this help.
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} { \
		if (/^[a-zA-Z_-]+:.*?##.*$$/) {printf "    ${YELLOW}%-20s${GREEN}%s${RESET}\n", $$1, $$2} \
		else if (/^## .*$$/) {printf "  ${CYAN}%s${RESET}\n", substr($$1,4)} \
		}' $(MAKEFILE_LIST)