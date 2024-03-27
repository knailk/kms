.PHONY: build

SRC_PATH:= ${PWD}

# ----- Setup Development Tools -----
prepare:
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.46.2
	@go install github.com/google/wire/cmd/wire@latest
	@go install golang.org/x/tools/cmd/goimports@latest
	@go install github.com/swaggo/swag/cmd/swag@latest

GO_MOD_ENV=GOPRIVATE=kms/*
mod:
	@$(GO_MOD_ENV) go mod tidy && go mod vendor

# ----- Build environment for developing -----
init: prepare mod docker-up

# ----- Development -----
build:
	@go build -o bin/application ${SRC_PATH}/cmd/server/...

run:
	@go run ${SRC_PATH}/cmd/server/...

local:
	ENV=local go run ${SRC_PATH}/cmd/server/...

dev:
	ENV=development go run ${SRC_PATH}/cmd/server/...

fmt: ## gofmt and goimports all go files
	@find . -name '*.go' -not -wholename './vendor/*' -not -wholename '*_gen.go' -not -wholename '*/mock_*.go' | while read -r file; do gofmt -w -s "$$file"; goimports -w "$$file"; done

lint:
	@golangci-lint --timeout=10m0s run -v ./...

check: mod lint test build

# ----- Code migration -----
mig-local-up:
	ENV=local go run ${SRC_PATH}/cmd/migration/... -up
mig-local-down:
	ENV=local go run ${SRC_PATH}/cmd/migration/... -down

# ----- Code generation -----
gen: gen-go gen-wire gen-repo

gen-go:
	@go generate ./...

gen-wire:
	@cd app/registry && wire

gen-repo:
	@go run app/external/persistence/database/gen/main.go

gen-swag: ## Swagger generate
	@swag fmt -d app/api/routes,app/api/handler -g routes.go && \
	swag init --ot go,yaml -d app/api/routes,./ -g routes.go -o app/api/docs --exclude pkg,db,deployments,scripts,vendor

# ----- Docker -----
docker-up:
	@docker-compose -p kms -f ${SRC_PATH}/development/docker-compose.dev.yml up -d

docker-down:
	@docker-compose -p kms -f ${SRC_PATH}/development/docker-compose.dev.yml down

docker-clear:
	@echo "Are you sure to remove volumes? [y/N]" && read ans && [ $${ans:-N} = y ]
	@docker-compose down -v
