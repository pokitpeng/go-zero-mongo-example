.PHONY: fmt go model run help

APP_NAME="go_zero_example"

build: ## build server
	bash ./scripts/build.sh ${APP_NAME}

fmt: ## format code
	goctl api format --dir ./api
	gofumpt -w ./api

go: fmt ## generate go code
	goctl api go -api ./api/*.api -dir . -style go_zero

model: ## generate model code
	goctl model mongo --type user,item --dir model/mongo -style go_zero -e --home tpl -c

swagger: fmt ## generate swagger docs
	rm -rf ./api/go_zero_example.json
	goctl api plugin -plugin goctl-swagger="swagger -filename go_zero_example.json" -api ./api/go_zero_example.api -dir ./api

run: ## run server
	go run .

help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help