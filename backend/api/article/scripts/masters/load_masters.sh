#!/bin/bash

set -e

GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

# Support both DB_PASSWORD and DB_PASS
[ -n "$DB_PASS" ] && [ -z "$DB_PASSWORD" ] && DB_PASSWORD="$DB_PASS"

export PGPASSWORD="$DB_PASSWORD"

echo -e "${GREEN}Loading master data...${NC}"

# Execute SQL file
execute_sql() {
    local file="$1"
    local desc="$2"

    if [ "$DB_HOST" = "localhost" ]; then
        docker compose cp "$file" postgres:/tmp/$(basename "$file")
        docker compose exec -T postgres psql -U "$DB_USER" -d "$DB_NAME" -c "SET search_path TO core; \i /tmp/$(basename "$file")" > /dev/null 2>&1
        docker compose exec -T postgres rm -f "/tmp/$(basename "$file")"
    else
        (echo "SET search_path TO core;"; cat "$file") | psql -h "$DB_HOST" -p "${DB_PORT:-5432}" -U "$DB_USER" -d "$DB_NAME" > /dev/null 2>&1
    fi

    echo -e "${GREEN}✓ $desc${NC}"
}

# Get count
get_count() {
    if [ "$DB_HOST" = "localhost" ]; then
        docker compose exec -T postgres psql -U "$DB_USER" -d "$DB_NAME" -c "SET search_path TO core; $1" -t 2>/dev/null | xargs
    else
        psql -h "$DB_HOST" -p "${DB_PORT:-5432}" -U "$DB_USER" -d "$DB_NAME" -c "SET search_path TO core; $1" -t 2>/dev/null | xargs
    fi
}

# Load master data files
for file in "$(dirname "$0")"/[0-9][0-9]_*.sql; do
    [ -f "$file" ] || continue
    desc=$(basename "$file" .sql | sed 's/^[0-9][0-9]_//' | tr '_-' ' ')
    execute_sql "$file" "$desc"
done

echo -e "${YELLOW}Summary: Categories: $(get_count "SELECT COUNT(*) FROM categories;"), Tags: $(get_count "SELECT COUNT(*) FROM tags;")${NC}"

unset PGPASSWORD
