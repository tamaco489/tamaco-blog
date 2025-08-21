#!/bin/bash

# Load master data into database
# Usage: ./scripts/masters/load_masters.sh

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Database connection settings
DB_HOST="localhost"
DB_PORT="5432"
DB_NAME="core"
DB_USER="core"
DB_PASSWORD="password#0"

# PostgreSQL connection string
PGPASSWORD="$DB_PASSWORD"
export PGPASSWORD

# Function to execute SQL file
execute_sql_file() {
    local sql_file="$1"
    local description="$2"

    echo -e "${YELLOW}Loading: $description${NC}"

    if [ ! -f "$sql_file" ]; then
        echo -e "${RED}Error: SQL file not found: $sql_file${NC}"
        return 1
    fi

    if docker compose exec -T postgres psql -U "$DB_USER" -h "$DB_HOST" -p "$DB_PORT" -d "$DB_NAME" -f "/tmp/$(basename "$sql_file")" > /dev/null 2>&1; then
        echo -e "${GREEN}✓ Successfully loaded: $description${NC}"
    else
        echo -e "${RED}✗ Failed to load: $description${NC}"
        return 1
    fi
}

# Function to copy SQL file to container and execute
load_sql_file() {
    local sql_file="$1"
    local description="$2"

    # Copy SQL file to container
    docker compose cp "$sql_file" postgres:/tmp/$(basename "$sql_file")

    # Execute SQL file
    execute_sql_file "$sql_file" "$description"

    # Clean up temporary file in container
    docker compose exec -T postgres rm -f "/tmp/$(basename "$sql_file")"
}

# Check if Docker containers are running
if ! docker compose ps postgres | grep -q "Up"; then
    echo -e "${RED}Error: PostgreSQL container is not running${NC}"
    echo "Please start the containers with: make up"
    exit 1
fi

echo -e "${GREEN}Starting master data loading...${NC}"
echo

# Directory containing master SQL files
MASTERS_DIR="$(dirname "$0")"

# Load master data files in order
if [ -f "$MASTERS_DIR/01_categories.sql" ]; then
    load_sql_file "$MASTERS_DIR/01_categories.sql" "Categories master data"
else
    echo -e "${YELLOW}Warning: Categories master data file not found${NC}"
fi

if [ -f "$MASTERS_DIR/02_tags.sql" ]; then
    load_sql_file "$MASTERS_DIR/02_tags.sql" "Tags master data"
else
    echo -e "${YELLOW}Warning: Tags master data file not found${NC}"
fi

# Load any additional master files (in alphabetical order)
for sql_file in "$MASTERS_DIR"/[0-9][0-9]_*.sql; do
    if [ -f "$sql_file" ]; then
        filename=$(basename "$sql_file")
        # Skip files we've already processed
        if [ "$filename" != "01_categories.sql" ] && [ "$filename" != "02_tags.sql" ]; then
            description=$(echo "$filename" | sed 's/[0-9][0-9]_\(.*\)\.sql/\1/' | tr '_-' ' ')
            load_sql_file "$sql_file" "$description master data"
        fi
    fi
done

echo
echo -e "${GREEN}✅ Master data loading completed successfully!${NC}"
echo

# Show summary
echo -e "${YELLOW}Summary:${NC}"
echo "Categories:"
docker compose exec -T postgres psql -U "$DB_USER" -h "$DB_HOST" -p "$DB_PORT" -d "$DB_NAME" -c "SELECT COUNT(*) as category_count FROM categories;" -t | xargs echo "  Total:"

echo "Tags:"
docker compose exec -T postgres psql -U "$DB_USER" -h "$DB_HOST" -p "$DB_PORT" -d "$DB_NAME" -c "SELECT COUNT(*) as tag_count FROM tags;" -t | xargs echo "  Total:"

unset PGPASSWORD
