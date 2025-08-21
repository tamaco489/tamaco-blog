#!/bin/bash

# Create migration file with sequential numbering
# Usage: ./scripts/create_migration.sh <migration_name>

set -e

# Check if migration name is provided
if [ -z "$1" ]; then
    echo "Error: Migration name is required"
    echo "Usage: $0 <migration_name>"
    echo "Example: $0 create_users_table"
    exit 1
fi

MIGRATION_NAME=$1
MIGRATION_DIR="internal/repository/migrations"

# Create migrations directory if it doesn't exist
mkdir -p "$MIGRATION_DIR"

# Get the next migration number
if [ -z "$(ls -A $MIGRATION_DIR 2>/dev/null)" ]; then
    # No migration files exist, start with 0001
    NEXT_NUMBER="0001"
else
    # Find the highest migration number and increment
    LAST_NUMBER=$(ls "$MIGRATION_DIR" | grep -E '^[0-9]{4}_' | sort -r | head -n1 | cut -d'_' -f1)

    if [ -z "$LAST_NUMBER" ]; then
        NEXT_NUMBER="0001"
    else
        # Remove leading zeros, increment, and pad back to 4 digits
        NEXT_NUMBER=$(printf "%04d" $((10#$LAST_NUMBER + 1)))
    fi
fi

# Create the migration file
MIGRATION_FILE="${MIGRATION_DIR}/${NEXT_NUMBER}_${MIGRATION_NAME}.sql"

# Check if file already exists
if [ -f "$MIGRATION_FILE" ]; then
    echo "Error: Migration file already exists: $MIGRATION_FILE"
    exit 1
fi

# Create migration file with template
cat > "$MIGRATION_FILE" << 'EOF'
-- +migrate Up

-- +migrate Down

EOF

echo "Created migration: $MIGRATION_FILE"
echo ""
echo "Next steps:"
echo "1. Edit the migration file to add your SQL statements"
echo "2. Run 'make migrate-up' to apply the migration"
echo "3. Run 'make migrate-status' to check migration status"
