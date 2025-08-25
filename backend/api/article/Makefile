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
	go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest
	go install github.com/rubenv/sql-migrate/...@latest
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	go install go.uber.org/mock/mockgen@latest
	npm i -g @redocly/cli@latest


# =================================================================
# generate open api spec and sqlc interfaces
# =================================================================
.PHONY: gen-api bundle-openapi gen-sqlc
gen-api: bundle-openapi ## dev: generate api server interface and type definitions based on the open api schema
	oapi-codegen -package gen -generate types -o internal/gen/types.gen.go ./spec/openapi.yaml
	oapi-codegen -package gen -generate strict-server,gin -templates ./spec/templates -o internal/gen/server.gen.go ./spec/openapi.yaml

bundle-openapi: ## dev: bundle open api spec
	redocly bundle ./spec/openapi_base.yaml --output ./spec/openapi.yaml

gen-sqlc: ## dev: generate sqlc interfaces and types
	sqlc generate -f internal/repository/sqlc.yaml
	mockgen -source=./internal/repository/gen_sqlc/querier.go -package mock -destination=./internal/repository/gen_mock/mock_querier.go
	mockgen -source=./internal/repository/gen_sqlc/db.go -package mock -destination=./internal/repository/gen_mock/mock_db.go

.PHONY: sqlc-lint sqlc-diff sqlc-vet
SQLC_DIRS := ./internal/repository

sqlc-lint: sqlc-diff sqlc-vet ## dev: run sqlc schema diff check and static analysis

sqlc-diff: ## dev: check sqlc schema differences
	@for dir in $(SQLC_DIRS); do echo "Running sqlc diff in $$dir"; cd $$dir && sqlc diff && cd - > /dev/null; done

sqlc-vet: ## dev: run sqlc static analysis
	@for dir in $(SQLC_DIRS); do echo "Running sqlc vet in $$dir"; cd $$dir && sqlc vet && cd - > /dev/null; done


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

migrate-new: ## dev: create a new migration with sequential numbering (e.g., make migrate-create NAME=create_users_table)
	@if [ -z "$(NAME)" ]; then \
		echo "Error: NAME is required"; \
		echo "Usage: make migrate-create NAME=<migration_name>"; \
		echo "Example: make migrate-create NAME=create_users_table"; \
		exit 1; \
	fi
	@./scripts/migration/create_migration.sh $(NAME)

migrate-up: ## dev: apply migration
	sql-migrate up -env='dev' -config=./_tools/sql-migrate/config.yaml

migrate-down: ## dev: rollback migration
	sql-migrate down -limit=1 -env='dev' -config=./_tools/sql-migrate/config.yaml

migrate-status: ## dev: show migration status
	sql-migrate status -env='dev' -config=./_tools/sql-migrate/config.yaml

.PHONY: supabase-migrate-up supabase-migrate-down supabase-migrate-status supabase-list-tables supabase-drop-tables
supabase-migrate-up: ## prd: apply migrations to database
	@SUPABASE_DB_URL="postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)" ./scripts/migration/supabase_migrate.sh up

supabase-migrate-down: ## prd: rollback last migration from database
	@SUPABASE_DB_URL="postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)" ./scripts/migration/supabase_migrate.sh down

supabase-migrate-status: ## prd: show migration status
	@SUPABASE_DB_URL="postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)" ./scripts/migration/supabase_migrate.sh status

supabase-list-tables: ## prd: list all tables in core schema
	@PGPASSWORD="$(DB_PASSWORD)" psql -h "$(DB_HOST)" -p "$(DB_PORT)" -U "$(DB_USER)" -d "$(DB_NAME)" -c "SELECT table_name FROM information_schema.tables WHERE table_schema = 'core' ORDER BY table_name;"

supabase-drop-tables: ## prd: drop all tables in core schema
	@PGPASSWORD="$(DB_PASSWORD)" psql -h "$(DB_HOST)" -p "$(DB_PORT)" -U "$(DB_USER)" -d "$(DB_NAME)" -c "DROP TABLE IF EXISTS core.article_seo_metadata, core.article_tags, core.articles, core.categories, core.tags, core.schema_migrations CASCADE;"

.PHONY: load-masters clear-masters init-masters
load-masters: ## dev/prd: load master data (usage: DB_HOST=localhost DB_PORT=5432 DB_NAME=core DB_USER=core DB_PASSWORD=password#0 make load-masters)
	@DB_HOST="$(DB_HOST)" DB_PORT="$(DB_PORT)" DB_NAME="$(DB_NAME)" DB_USER="$(DB_USER)" DB_PASSWORD="$(DB_PASSWORD)" ./scripts/masters/load_masters.sh

clear-masters: ## dev/prd: clear all master data (usage: DB_HOST=localhost DB_PORT=5432 DB_NAME=core DB_USER=core DB_PASSWORD=password#0 make clear-masters)
	@DB_HOST="$(DB_HOST)" DB_PORT="$(DB_PORT)" DB_NAME="$(DB_NAME)" DB_USER="$(DB_USER)" DB_PASSWORD="$(DB_PASSWORD)" ./scripts/masters/clear_masters.sh

init-masters: ## dev/prd: reset and reload master data (usage: DB_HOST=localhost DB_PORT=5432 DB_NAME=core DB_USER=core DB_PASSWORD=password#0 make init-masters)
	@$(MAKE) clear-masters DB_HOST="$(DB_HOST)" DB_PORT="$(DB_PORT)" DB_NAME="$(DB_NAME)" DB_USER="$(DB_USER)" DB_PASSWORD="$(DB_PASSWORD)"
	@$(MAKE) load-masters DB_HOST="$(DB_HOST)" DB_PORT="$(DB_PORT)" DB_NAME="$(DB_NAME)" DB_USER="$(DB_USER)" DB_PASSWORD="$(DB_PASSWORD)"


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

# ecr
VERSION := article_api_v0.0.0
.PHONY: auth push get-images delete-images
auth: ## prd: authenticate ecr
	aws ecr get-login-password --region ap-northeast-1 --profile ${AWS_PROFILE} | docker login --username AWS --password-stdin ${AWS_ACCOUNT_ID}.dkr.ecr.ap-northeast-1.amazonaws.com

push: auth ## prd: push image to ecr
	docker build --platform linux/amd64 --no-cache -t $(AWS_ACCOUNT_ID).dkr.ecr.ap-northeast-1.amazonaws.com/$(ENV)-article-api:$(VERSION) -f ./container/Dockerfile.article .
	docker push ${AWS_ACCOUNT_ID}.dkr.ecr.ap-northeast-1.amazonaws.com/$(ENV)-article-api:$(VERSION)

get-images: ## prd: list images in ecr
	aws ecr list-images --repository-name $(ENV)-article-api --region ap-northeast-1 --profile ${AWS_PROFILE} --query 'imageIds[*]' | jq .

delete-images: ## prd: delete images in ecr
	@aws ecr list-images --repository-name $(ENV)-article-api --region ap-northeast-1 --profile ${AWS_PROFILE} --query 'imageIds[*]' --output json > /tmp/images.json
	@cat /tmp/images.json | jq -c '.[]' | while read -r image_id; do \
		aws ecr batch-delete-image --repository-name $(ENV)-article-api --region ap-northeast-1 --profile ${AWS_PROFILE} --image-ids $${image_id}; \
	done | jq .

# lambda
.PHONY: deploy
deploy: push ## prd: deploy lambda function
	aws lambda update-function-code  --profile ${AWS_PROFILE} \
		--function-name $(ENV)-tamaco-blog-article-api \
		--image-uri $(AWS_ACCOUNT_ID).dkr.ecr.ap-northeast-1.amazonaws.com/$(ENV)-article-api:$(VERSION) | jq .


# =================================================================
# api request
# =================================================================
.PHONY: health
health: ## dev: check health of api server
	curl -sX 'GET' \
		'http://localhost:8080/v1/health' \
		-H 'accept: application/json' | jq .

get-articles: ## dev: get articles
	curl -sX 'GET' \
		'http://localhost:8080/v1/articles' \
		-H 'accept: application/json' | jq .

get-categories: ## dev: get categories
	curl -sX 'GET' \
		'http://localhost:8080/v1/categories' \
		-H 'accept: application/json' | jq .

get-tags: ## dev: get tags
	curl -sX 'GET' \
		'http://localhost:8080/v1/tags' \
		-H 'accept: application/json' | jq .


# =================================================================
# other
# =================================================================
.PHONY: help
help: ## Help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
