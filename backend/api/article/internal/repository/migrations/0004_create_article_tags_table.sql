-- +migrate Up
CREATE TABLE IF NOT EXISTS core.article_tags (
    "article_id" UUID NOT NULL REFERENCES core.articles("id") ON DELETE CASCADE,
    "tag_id" UUID NOT NULL REFERENCES core.tags("id") ON DELETE CASCADE,
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("article_id", "tag_id")
);

CREATE INDEX IF NOT EXISTS idx_article_tags_article_id ON core.article_tags("article_id");
CREATE INDEX IF NOT EXISTS idx_article_tags_tag_id ON core.article_tags("tag_id");

COMMENT ON TABLE core.article_tags IS '記事とタグの多対多リレーション';
COMMENT ON COLUMN core.article_tags."article_id" IS '記事ID';
COMMENT ON COLUMN core.article_tags."tag_id" IS 'タグID';

-- +migrate Down
DROP TABLE IF EXISTS core.article_tags CASCADE;
