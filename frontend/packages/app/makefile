
.PHONY: help, build

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

run: install ## Start development server
	@pnpm dev

build: install ## Create production build
	@pnpm build

publish: build ## Publish the frontend
	zip -r build.zip ./build/*
	scp build.zip poker:/opt/poker

install: ## Install dependencies
	@pnpm install

preview: ## Preview the production build locally
	http-server ./build -d -i