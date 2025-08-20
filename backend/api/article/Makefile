# =================================================================
# tools for setting up the execution environment
# =================================================================
.PHONY: setup-env install-tools
setup-env: ## dev: Set environment variables
	cp -p .env_sample .env
	cp -p .env_localstack_sample .env_localstack

install-tools: setup-env ## dev: Install tools used in the project
	go install github.com/air-verse/air@latest
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
	go install golang.org/x/tools/cmd/godoc@latest
	go install golang.org/x/tools/cmd/goimports@latest
	go install golang.org/x/vuln/cmd/govulncheck@latest
	go install github.com/rubenv/sql-migrate/...@latest
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	go install go.uber.org/mock/mockgen@latest
	npm i -g @redocly/cli@latest


# =================================================================
# generate open api spec and sqlc interfaces
# =================================================================
.PHONY: gen-api bundle-openapi
gen-api: bundle-openapi ## dev: generate api server interface and type definitions based on the open api schema
	oapi-codegen -package gen -generate types -o internal/gen/types.gen.go ./spec/openapi.yaml
	oapi-codegen -package gen -generate strict-server,gin -templates ./spec/templates -o internal/gen/server.gen.go ./spec/openapi.yaml

bundle-openapi: ## dev: bundle open api spec
	redocly bundle ./spec/openapi_base.yaml --output ./spec/openapi.yaml

gen-sqlc: ## dev: generate sqlc interfaces and types
	sqlc generate -f internal/repository/sqlc.yaml
	mockgen -source=./internal/repository/gen_sqlc/querier.go -package mock -destination=./internal/repository/gen_mock/mock_querier.go
	mockgen -source=./internal/repository/gen_sqlc/db.go -package mock -destination=./internal/repository/gen_mock/mock_db.go


# =================================================================
# local execution environment setup
# =================================================================
.PHONY: lint fmt vuln deps build run
lint: ## dev: static code analysis
	golangci-lint run --timeout 3m

fmt: ## dev: format go code
	goimports -w .

vuln: ## dev: check for security vulnerabilities
	govulncheck ./...

deps: ## dev: download dependencies
	go mod download && go mod tidy

build: deps ## dev: build the project
	go build -o build/article_api cmd/main.go

run: build ## dev: run api server
	./build/article_api


# =================================================================
# container execution environment setup
# =================================================================
.PHONY: up down logs
up: ## dev: start containers
	docker compose up -d api postgres localstack

down: ## dev: stop and remove containers
	docker compose down api postgres localstack

logs: ## dev: show logs
	docker compose logs -f api

rebuild: ## dev: rebuild containers
	docker compose down -v
	docker compose build --no-cache
	docker compose up -d api postgres localstack

# =================================================================
# datastore operations
# =================================================================
.PHONY: pg
pg: ## dev: connect to postgres
	PGPASSWORD="password#0" docker compose exec postgres psql -U core -h localhost -p 5432 -d core

.PHONY: migrate-new migrate-up migrate-down migrate-status
migrate-new: ## dev: create a new migration schema
	sql-migrate new -env='dev' -config=./_tools/sql-migrate/config.yaml $(NAME)

migrate-up: ## dev: apply migration
	sql-migrate up -env='dev' -config=./_tools/sql-migrate/config.yaml

migrate-down: ## dev: rollback migration
	sql-migrate down -limit=1 -env='dev' -config=./_tools/sql-migrate/config.yaml

migrate-status: ## dev: show migration status
	sql-migrate status -env='dev' -config=./_tools/sql-migrate/config.yaml


# =================================================================
# aws operations
# =================================================================
# secrets manager
.PHONY: get-secrets get-secret-value
get-secrets: ## dev: get all secrets from secrets manager
	docker exec -it article_localstack awslocal secretsmanager list-secrets --region ap-northeast-1

get-secret-value: ## dev: get secret value from secrets manager
	docker exec -it article_localstack awslocal secretsmanager get-secret-value \
		--secret-id dev/tamaco-blog/article/core/rds \
		--region ap-northeast-1


# =================================================================
# api request
# =================================================================
.PHONY: health
health: ## dev: check health of api server
	curl -sX 'GET' \
		'http://localhost:8080/v1/health' \
		-H 'accept: application/json' | jq .


# =================================================================
# other
# =================================================================
.PHONY: help
help: ## Help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
