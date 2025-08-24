-- +migrate Up
CREATE TABLE IF NOT EXISTS core.categories (
    "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "name" VARCHAR(100) NOT NULL,
    "slug" VARCHAR(100) NOT NULL UNIQUE,
    "description" TEXT,
    "display_order" INTEGER DEFAULT 0,
    "article_count" INTEGER DEFAULT 0,
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_categories_slug ON core.categories("slug");

COMMENT ON TABLE core.categories IS 'ブログ記事のカテゴリー';
COMMENT ON COLUMN core.categories."id" IS 'カテゴリの一意識別子';
COMMENT ON COLUMN core.categories."name" IS 'カテゴリ名';
COMMENT ON COLUMN core.categories."slug" IS 'URL用の識別子';
COMMENT ON COLUMN core.categories."description" IS 'カテゴリの説明';
COMMENT ON COLUMN core.categories."display_order" IS '表示順';
COMMENT ON COLUMN core.categories."article_count" IS 'このカテゴリの記事数';

-- +migrate Down
DROP TABLE IF EXISTS core.categories CASCADE;
