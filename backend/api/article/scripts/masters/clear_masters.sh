#!/bin/bash

set -e

GREEN='\033[0;32m'
NC='\033[0m'

# Support both DB_PASSWORD and DB_PASS
[ -n "$DB_PASS" ] && [ -z "$DB_PASSWORD" ] && DB_PASSWORD="$DB_PASS"

export PGPASSWORD="$DB_PASSWORD"

echo -e "${GREEN}Clearing master data...${NC}"

if [ "$DB_HOST" = "localhost" ]; then
    docker compose exec -T postgres psql -U "$DB_USER" -d "$DB_NAME" -c "TRUNCATE TABLE article_tags, articles, tags, categories RESTART IDENTITY CASCADE;"
else
    psql -h "$DB_HOST" -p "${DB_PORT:-5432}" -U "$DB_USER" -d "$DB_NAME" -c "TRUNCATE TABLE article_tags, articles, tags, categories RESTART IDENTITY CASCADE;"
fi

echo -e "${GREEN}✓ Master data cleared${NC}"
unset PGPASSWORD
