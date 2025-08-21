-- name: GetArticle :one
SELECT * FROM "articles"
WHERE "id" = $1;

-- name: GetArticleBySlug :one
SELECT * FROM "articles"
WHERE "slug" = $1;

-- name: ListArticles :many
SELECT * FROM "articles"
WHERE "status" = 'published'
ORDER BY "published_at" DESC
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

-- name: CountArticles :one
SELECT COUNT(*) FROM "articles"
WHERE "status" = 'published';

-- name: ListArticlesByStatus :many
SELECT * FROM "articles"
WHERE "status" = sqlc.arg('status')
ORDER BY "created_at" DESC
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

-- name: ListArticlesByCategory :many
SELECT * FROM "articles"
WHERE "category_id" = sqlc.arg('category_id') AND "status" = 'published'
ORDER BY "published_at" DESC
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

-- name: CountArticlesByCategory :one
SELECT COUNT(*) FROM "articles"
WHERE "category_id" = sqlc.arg('category_id') AND "status" = 'published';

-- name: ListRecentArticles :many
SELECT * FROM "articles"
WHERE "status" = 'published'
ORDER BY "published_at" DESC
LIMIT sqlc.arg('limit');

-- name: ListPopularArticles :many
SELECT * FROM "articles"
WHERE "status" = 'published' AND "view_count" > 0
ORDER BY "view_count" DESC
LIMIT sqlc.arg('limit');

-- name: ListArticlesByMonth :many
SELECT * FROM "articles"
WHERE "status" = 'published'
  AND EXTRACT(YEAR FROM "published_at") = sqlc.arg('year')
  AND EXTRACT(MONTH FROM "published_at") = sqlc.arg('month')
ORDER BY "published_at" DESC
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

-- name: CountArticlesByMonth :one
SELECT COUNT(*) FROM "articles"
WHERE "status" = 'published'
  AND EXTRACT(YEAR FROM "published_at") = sqlc.arg('year')
  AND EXTRACT(MONTH FROM "published_at") = sqlc.arg('month');

-- name: SearchArticles :many
SELECT * FROM "articles"
WHERE "status" = 'published'
  AND (
    "title" ILIKE '%' || sqlc.arg('query') || '%' OR
    "content" ILIKE '%' || sqlc.arg('query') || '%' OR
    "summary" ILIKE '%' || sqlc.arg('query') || '%'
  )
ORDER BY "published_at" DESC
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

-- name: CountSearchArticles :one
SELECT COUNT(*) FROM "articles"
WHERE "status" = 'published'
  AND (
    "title" ILIKE '%' || sqlc.arg('query') || '%' OR
    "content" ILIKE '%' || sqlc.arg('query') || '%' OR
    "summary" ILIKE '%' || sqlc.arg('query') || '%'
  );

-- name: CreateArticle :one
INSERT INTO "articles" (
    "title", "content", "slug", "summary", "status", "category_id", "published_at"
) VALUES (
    sqlc.arg('title'), sqlc.arg('content'), sqlc.arg('slug'),
    sqlc.narg('summary'), sqlc.arg('status'), sqlc.narg('category_id'),
    sqlc.narg('published_at')
) RETURNING *;

-- name: UpdateArticle :one
UPDATE "articles"
SET
    "title" = sqlc.arg('title'),
    "content" = sqlc.arg('content'),
    "slug" = sqlc.arg('slug'),
    "summary" = sqlc.narg('summary'),
    "status" = sqlc.arg('status'),
    "category_id" = sqlc.narg('category_id'),
    "published_at" = sqlc.narg('published_at'),
    "updated_at" = CURRENT_TIMESTAMP
WHERE "id" = sqlc.arg('id')
RETURNING *;

-- name: UpdateArticleStatus :one
UPDATE "articles"
SET
    "status" = sqlc.arg('status'),
    "published_at" = CASE
        WHEN sqlc.arg('status') = 'published' AND "published_at" IS NULL THEN CURRENT_TIMESTAMP
        ELSE "published_at"
    END,
    "updated_at" = CURRENT_TIMESTAMP
WHERE "id" = sqlc.arg('id')
RETURNING *;

-- name: IncrementArticleViewCount :exec
UPDATE "articles"
SET "view_count" = "view_count" + 1
WHERE "id" = $1;

-- name: DeleteArticle :exec
DELETE FROM "articles"
WHERE "id" = $1;
