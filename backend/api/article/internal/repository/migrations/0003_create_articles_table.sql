-- +migrate Up
CREATE TABLE IF NOT EXISTS core.articles (
    "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "title" VARCHAR(200) NOT NULL,
    "content" TEXT NOT NULL,
    "slug" VARCHAR(200) NOT NULL UNIQUE,
    "summary" VARCHAR(500),
    "status" VARCHAR(20) NOT NULL DEFAULT 'draft' CHECK ("status" IN ('draft', 'published', 'private')),
    "category_id" UUID REFERENCES core.categories("id") ON DELETE SET NULL,
    "view_count" INTEGER DEFAULT 0,
    "published_at" TIMESTAMP WITH TIME ZONE,
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_articles_slug ON core.articles("slug");
CREATE INDEX idx_articles_status ON core.articles("status");
CREATE INDEX idx_articles_published_at ON core.articles("published_at" DESC);

COMMENT ON TABLE core.articles IS 'ブログ記事';
COMMENT ON COLUMN core.articles."id" IS '記事の一意識別子';
COMMENT ON COLUMN core.articles."title" IS '記事のタイトル';
COMMENT ON COLUMN core.articles."content" IS 'Markdown形式の記事本文';
COMMENT ON COLUMN core.articles."slug" IS 'URL用の識別子';
COMMENT ON COLUMN core.articles."summary" IS '記事の要約（SEO用）';
COMMENT ON COLUMN core.articles."status" IS '記事の公開状態: draft(下書き), published(公開), private(非公開)';
COMMENT ON COLUMN core.articles."category_id" IS 'カテゴリID';
COMMENT ON COLUMN core.articles."view_count" IS '閲覧数';
COMMENT ON COLUMN core.articles."published_at" IS '公開日時';

-- +migrate Down
DROP TABLE IF EXISTS core.articles CASCADE;
