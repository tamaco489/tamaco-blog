-- name: GetArticleSEOMetadata :one
SELECT * FROM "article_seo_metadata"
WHERE "article_id" = $1;

-- name: CreateArticleSEOMetadata :one
INSERT INTO "article_seo_metadata" (
    "article_id",
    "description",
    "keywords",
    "og_image"
) VALUES (
    sqlc.arg('article_id'),
    sqlc.narg('description'),
    sqlc.narg('keywords'),
    sqlc.narg('og_image')
) RETURNING *;

-- name: UpdateArticleSEOMetadata :one
UPDATE "article_seo_metadata"
SET
    "description" = sqlc.narg('description'),
    "keywords" = sqlc.narg('keywords'),
    "og_image" = sqlc.narg('og_image'),
    "updated_at" = CURRENT_TIMESTAMP
WHERE "article_id" = sqlc.arg('article_id')
RETURNING *;

-- name: UpsertArticleSEOMetadata :one
INSERT INTO "article_seo_metadata" (
    "article_id",
    "description",
    "keywords",
    "og_image"
) VALUES (
    sqlc.arg('article_id'),
    sqlc.narg('description'),
    sqlc.narg('keywords'),
    sqlc.narg('og_image')
)
ON CONFLICT ("article_id") DO UPDATE
SET
    "description" = EXCLUDED."description",
    "keywords" = EXCLUDED."keywords",
    "og_image" = EXCLUDED."og_image",
    "updated_at" = CURRENT_TIMESTAMP
RETURNING *;

-- name: DeleteArticleSEOMetadata :exec
DELETE FROM "article_seo_metadata"
WHERE "article_id" = $1;

-- name: ListArticlesWithSEOMetadata :many
SELECT
    "a".*,
    "sm"."description" AS "seo_description",
    "sm"."keywords" AS "seo_keywords",
    "sm"."og_image" AS "seo_og_image"
FROM "articles" "a"
LEFT JOIN "article_seo_metadata" "sm" ON "a"."id" = "sm"."article_id"
WHERE "a"."status" = 'published'
ORDER BY "a"."published_at" DESC
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');
