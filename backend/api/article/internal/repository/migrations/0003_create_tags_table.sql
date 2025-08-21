-- +migrate Up
CREATE TABLE IF NOT EXISTS "tags" (
    "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    "name" VARCHAR(50) NOT NULL,
    "slug" VARCHAR(50) NOT NULL UNIQUE,
    "usage_count" INTEGER DEFAULT 0,
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX "idx_tags_slug" ON "tags"("slug");

COMMENT ON TABLE "tags" IS 'ブログ記事のタグ';
COMMENT ON COLUMN "tags"."id" IS 'タグの一意識別子';
COMMENT ON COLUMN "tags"."name" IS 'タグ名';
COMMENT ON COLUMN "tags"."slug" IS 'URL用の識別子';
COMMENT ON COLUMN "tags"."usage_count" IS 'タグが使用されている記事数';

-- +migrate Down
DROP TABLE IF EXISTS "tags";
