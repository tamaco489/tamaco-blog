-- name: GetCategory :one
SELECT * FROM "categories"
WHERE "id" = $1;

-- name: GetCategoryBySlug :one
SELECT * FROM "categories"
WHERE "slug" = $1;

-- name: ListCategories :many
SELECT * FROM "categories"
ORDER BY "display_order" ASC, "name" ASC;

-- name: CreateCategory :one
INSERT INTO "categories" (
    "name",
    "slug",
    "description",
    "display_order"
) VALUES (
    sqlc.arg('name'),
    sqlc.arg('slug'),
    sqlc.narg('description'),
    sqlc.arg('display_order')
) RETURNING *;

-- name: UpdateCategory :one
UPDATE "categories"
SET
    "name" = sqlc.arg('name'),
    "slug" = sqlc.arg('slug'),
    "description" = sqlc.narg('description'),
    "display_order" = sqlc.arg('display_order'),
    "updated_at" = CURRENT_TIMESTAMP
WHERE "id" = sqlc.arg('id')
RETURNING *;

-- name: DeleteCategory :exec
DELETE FROM "categories"
WHERE "id" = $1;

-- name: UpdateCategoryArticleCount :exec
UPDATE "categories"
SET
    "article_count" = sqlc.arg('article_count'),
    "updated_at" = CURRENT_TIMESTAMP
WHERE "id" = sqlc.arg('id');
