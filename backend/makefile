.PHONY: help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

run: ## Run the service in dev mode
	# @go run ./src/
	go run github.com/cortesi/modd/cmd/modd --file=modd-run.conf

build: ## Build for linux
	env GOOS=linux GOARCH=amd64 go build -o poker ./src/

buildlocal: ## Build for local
	go build -o poker ./src/

buildandrun: buildlocal ## _
	./poker


publish: build ## Build and upload
	scp ./poker poker:/opt/poker
