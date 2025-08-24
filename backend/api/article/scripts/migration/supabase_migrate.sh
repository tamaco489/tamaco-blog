#!/bin/bash

# Supabaseデータベースマイグレーションスクリプト
#
# 使用方法:
#   ./supabase_migrate.sh up    # マイグレーション適用
#   ./supabase_migrate.sh status # 状態確認
#   ./supabase_migrate.sh down   # ロールバック（1つ前に戻る）

set -e

# 色付きメッセージ
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# 設定
MIGRATION_DIR="./internal/repository/migrations"
MIGRATION_TABLE="schema_migrations"

# 環境変数から接続情報を取得
if [ -z "$SUPABASE_DB_URL" ]; then
    echo -e "${RED}Error: SUPABASE_DB_URL environment variable is not set${NC}"
    echo "Example: postgresql://postgres.[project-ref]:[password]@aws-0-ap-northeast-1.pooler.supabase.com:5432/postgres"
    exit 1
fi

# マイグレーションテーブルの作成
create_migration_table() {
    echo -e "${YELLOW}Creating migration table if not exists...${NC}"
    psql "$SUPABASE_DB_URL" <<-EOF
        CREATE TABLE IF NOT EXISTS ${MIGRATION_TABLE} (
            id SERIAL PRIMARY KEY,
            version VARCHAR(255) NOT NULL UNIQUE,
            applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        );
EOF
}

# 適用済みマイグレーションの取得
get_applied_migrations() {
    psql "$SUPABASE_DB_URL" -t -A -c "SELECT version FROM ${MIGRATION_TABLE} ORDER BY version;"
}

# マイグレーションの適用
apply_migration() {
    local file=$1
    local version=$(basename "$file" .sql)

    echo -e "${YELLOW}Applying migration: $version${NC}"

    # トランザクション内でマイグレーションを実行
    psql "$SUPABASE_DB_URL" <<-EOF
        BEGIN;
        -- Set search path to use core schema
        SET search_path TO core, public;

        -- マイグレーションファイルのUpセクションのみを実行
        $(sed -n '/-- +migrate Up/,/-- +migrate Down/p' "$file" | sed '$ d')

        -- マイグレーション履歴に記録
        INSERT INTO ${MIGRATION_TABLE} (version) VALUES ('$version');

        COMMIT;
EOF

    if [ $? -eq 0 ]; then
        echo -e "${GREEN}✓ Applied: $version${NC}"
    else
        echo -e "${RED}✗ Failed: $version${NC}"
        exit 1
    fi
}

# マイグレーションのロールバック
rollback_migration() {
    local last_version=$(psql "$SUPABASE_DB_URL" -t -A -c "SELECT version FROM ${MIGRATION_TABLE} ORDER BY version DESC LIMIT 1;")

    if [ -z "$last_version" ]; then
        echo -e "${YELLOW}No migrations to rollback${NC}"
        return
    fi

    echo -e "${YELLOW}Rolling back migration: $last_version${NC}"

    # マイグレーションファイルのDownセクションを実行
    local migration_file="${MIGRATION_DIR}/${last_version}.sql"
    if [ -f "$migration_file" ]; then
        # Downセクションが存在するかチェック
        if grep -q "\-\- +migrate Down" "$migration_file"; then
            psql "$SUPABASE_DB_URL" <<-EOF
                BEGIN;
                -- Set search path to use core schema
                SET search_path TO core, public;

                -- マイグレーションファイルのDownセクションのみを実行
                $(sed -n '/-- +migrate Down/,$ p' "$migration_file" | tail -n +2)

                -- マイグレーション履歴から削除
                DELETE FROM ${MIGRATION_TABLE} WHERE version = '$last_version';

                COMMIT;
EOF
            if [ $? -eq 0 ]; then
                echo -e "${GREEN}✓ Rolled back: $last_version${NC}"
            else
                echo -e "${RED}✗ Failed to rollback: $last_version${NC}"
                exit 1
            fi
        else
            echo -e "${YELLOW}Warning: No rollback section found for $last_version${NC}"
            echo "Only removing from migration history"

            # 履歴からは削除
            psql "$SUPABASE_DB_URL" -c "DELETE FROM ${MIGRATION_TABLE} WHERE version = '$last_version';"
            echo -e "${GREEN}✓ Removed from history: $last_version${NC}"
        fi
    else
        echo -e "${RED}Error: Migration file not found: $migration_file${NC}"
        exit 1
    fi
}

# ステータス表示
show_status() {
    echo -e "${YELLOW}Migration Status:${NC}"
    echo "-------------------"

    # 適用済みマイグレーション
    echo -e "${GREEN}Applied migrations:${NC}"
    local applied=$(get_applied_migrations)
    if [ -z "$applied" ]; then
        echo "  (none)"
    else
        echo "$applied" | while read -r version; do
            echo "  ✓ $version"
        done
    fi

    echo ""

    # 未適用マイグレーション
    echo -e "${YELLOW}Pending migrations:${NC}"
    local has_pending=false
    for file in $(ls -1 ${MIGRATION_DIR}/*.sql 2>/dev/null | grep -v ".down.sql" | sort); do
        local version=$(basename "$file" .sql)
        if ! echo "$applied" | grep -q "^$version$"; then
            echo "  ○ $version"
            has_pending=true
        fi
    done

    if [ "$has_pending" = false ]; then
        echo "  (none)"
    fi
}

# マイグレーションアップ
migrate_up() {
    create_migration_table

    local applied=$(get_applied_migrations)
    local has_new=false

    for file in $(ls -1 ${MIGRATION_DIR}/*.sql 2>/dev/null | grep -v ".down.sql" | sort); do
        local version=$(basename "$file" .sql)

        # 既に適用済みか確認
        if echo "$applied" | grep -q "^$version$"; then
            continue
        fi

        has_new=true
        apply_migration "$file"
    done

    if [ "$has_new" = false ]; then
        echo -e "${GREEN}All migrations are already applied${NC}"
    fi
}

# メイン処理
main() {
    case "$1" in
        up)
            migrate_up
            ;;
        down)
            create_migration_table
            rollback_migration
            ;;
        status)
            create_migration_table
            show_status
            ;;
        *)
            echo "Usage: $0 {up|down|status}"
            echo ""
            echo "Commands:"
            echo "  up     - Apply all pending migrations"
            echo "  down   - Rollback the last migration"
            echo "  status - Show migration status"
            exit 1
            ;;
    esac
}

main "$@"
