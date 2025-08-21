-- name: AddArticleTag :exec
INSERT INTO "article_tags" (
    "article_id",
    "tag_id"
) VALUES (
    sqlc.arg('article_id'),
    sqlc.arg('tag_id')
) ON CONFLICT ("article_id", "tag_id") DO NOTHING;

-- name: RemoveArticleTag :exec
DELETE FROM "article_tags"
WHERE "article_id" = sqlc.arg('article_id') AND "tag_id" = sqlc.arg('tag_id');

-- name: RemoveAllArticleTags :exec
DELETE FROM "article_tags"
WHERE "article_id" = $1;

-- name: ListTagsByArticleID :many
SELECT "t".* FROM "tags" "t"
INNER JOIN "article_tags" "at" ON "t"."id" = "at"."tag_id"
WHERE "at"."article_id" = $1
ORDER BY "t"."name" ASC;

-- name: ListArticlesByTagID :many
SELECT "a".* FROM "articles" "a"
INNER JOIN "article_tags" "at" ON "a"."id" = "at"."article_id"
WHERE "at"."tag_id" = sqlc.arg('tag_id') AND "a"."status" = 'published'
ORDER BY "a"."published_at" DESC
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

-- name: CountArticlesByTagID :one
SELECT COUNT(*) FROM "articles" "a"
INNER JOIN "article_tags" "at" ON "a"."id" = "at"."article_id"
WHERE "at"."tag_id" = sqlc.arg('tag_id') AND "a"."status" = 'published';

-- name: ListArticleIDsByTagID :many
SELECT "article_id" FROM "article_tags"
WHERE "tag_id" = $1;

-- name: CheckArticleHasTag :one
SELECT EXISTS(
    SELECT 1 FROM "article_tags"
    WHERE "article_id" = sqlc.arg('article_id') AND "tag_id" = sqlc.arg('tag_id')
);

-- name: AddArticleTagBulk :exec
INSERT INTO "article_tags" (
    "article_id",
    "tag_id"
) VALUES (
    sqlc.arg('article_id'),
    sqlc.arg('tag_id')
) ON CONFLICT ("article_id", "tag_id") DO NOTHING;
