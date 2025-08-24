-- +migrate Up
CREATE TABLE IF NOT EXISTS core.tags (
    "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "name" VARCHAR(50) NOT NULL,
    "slug" VARCHAR(50) NOT NULL UNIQUE,
    "usage_count" INTEGER DEFAULT 0,
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_tags_slug ON core.tags("slug");

COMMENT ON TABLE core.tags IS 'ブログ記事のタグ';
COMMENT ON COLUMN core.tags."id" IS 'タグの一意識別子';
COMMENT ON COLUMN core.tags."name" IS 'タグ名';
COMMENT ON COLUMN core.tags."slug" IS 'URL用の識別子';
COMMENT ON COLUMN core.tags."usage_count" IS 'タグが使用されている記事数';

-- +migrate Down
DROP TABLE IF EXISTS core.tags CASCADE;
