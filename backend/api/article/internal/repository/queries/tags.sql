-- name: GetTag :one
SELECT * FROM "tags"
WHERE "id" = $1;

-- name: GetTagBySlug :one
SELECT * FROM "tags"
WHERE "slug" = $1;

-- name: ListTags :many
SELECT * FROM "tags"
ORDER BY "usage_count" DESC, "name" ASC;

-- name: ListPopularTags :many
SELECT * FROM "tags"
WHERE "usage_count" > 0
ORDER BY "usage_count" DESC
LIMIT sqlc.arg('limit');

-- name: CreateTag :one
INSERT INTO "tags" (
    "name", "slug"
) VALUES (
    sqlc.arg('name'), sqlc.arg('slug')
) RETURNING *;

-- name: UpdateTag :one
UPDATE "tags"
SET
    "name" = sqlc.arg('name'),
    "slug" = sqlc.arg('slug'),
    "updated_at" = CURRENT_TIMESTAMP
WHERE "id" = sqlc.arg('id')
RETURNING *;

-- name: DeleteTag :exec
DELETE FROM "tags"
WHERE "id" = $1;

-- name: UpdateTagUsageCount :exec
UPDATE "tags"
SET
    "usage_count" = sqlc.arg('usage_count'),
    "updated_at" = CURRENT_TIMESTAMP
WHERE "id" = sqlc.arg('id');

-- name: IncrementTagUsageCount :exec
UPDATE "tags"
SET
    "usage_count" = "usage_count" + 1,
    "updated_at" = CURRENT_TIMESTAMP
WHERE "id" = $1;

-- name: DecrementTagUsageCount :exec
UPDATE "tags"
SET
    "usage_count" = GREATEST("usage_count" - 1, 0),
    "updated_at" = CURRENT_TIMESTAMP
WHERE "id" = $1;
