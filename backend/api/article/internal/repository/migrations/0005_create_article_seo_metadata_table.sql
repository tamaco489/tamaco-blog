-- +migrate Up
CREATE TABLE IF NOT EXISTS article_seo_metadata (
    "article_id" UUID PRIMARY KEY REFERENCES articles("id") ON DELETE CASCADE,
    "description" VARCHAR(160),
    "keywords" TEXT[], -- PostgreSQL array type for storing multiple keywords
    "og_image" VARCHAR(500),
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

COMMENT ON TABLE article_seo_metadata IS '記事のSEOメタデータ';
COMMENT ON COLUMN article_seo_metadata."article_id" IS '記事ID';
COMMENT ON COLUMN article_seo_metadata."description" IS 'メタディスクリプション';
COMMENT ON COLUMN article_seo_metadata."keywords" IS 'SEO用キーワード配列';
COMMENT ON COLUMN article_seo_metadata."og_image" IS 'OGP画像URL';

-- +migrate Down
DROP TABLE IF EXISTS article_seo_metadata;
