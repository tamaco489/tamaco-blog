-- name: GetArticleWithRelations :one
SELECT
    "a".*,
    "c"."name" AS "category_name",
    "c"."slug" AS "category_slug",
    "sm"."description" AS "seo_description",
    "sm"."keywords" AS "seo_keywords",
    "sm"."og_image" AS "seo_og_image"
FROM "articles" "a"
LEFT JOIN "categories" "c" ON "a"."category_id" = "c"."id"
LEFT JOIN "article_seo_metadata" "sm" ON "a"."id" = "sm"."article_id"
WHERE "a"."id" = $1;

-- name: GetArticleBySlugWithRelations :one
SELECT
    "a".*,
    "c"."name" AS "category_name",
    "c"."slug" AS "category_slug",
    "sm"."description" AS "seo_description",
    "sm"."keywords" AS "seo_keywords",
    "sm"."og_image" AS "seo_og_image"
FROM "articles" "a"
LEFT JOIN "categories" "c" ON "a"."category_id" = "c"."id"
LEFT JOIN "article_seo_metadata" "sm" ON "a"."id" = "sm"."article_id"
WHERE "a"."slug" = $1;

-- name: ListArticlesWithCategory :many
SELECT
    "a".*,
    "c"."name" AS "category_name",
    "c"."slug" AS "category_slug"
FROM "articles" "a"
LEFT JOIN "categories" "c" ON "a"."category_id" = "c"."id"
WHERE "a"."status" = 'published'
ORDER BY "a"."published_at" DESC
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

-- name: GetArticleWithTags :many
SELECT
    "a".*,
    COALESCE(
        json_agg(
            json_build_object(
                'id', "t"."id",
                'name', "t"."name",
                'slug', "t"."slug"
            ) ORDER BY "t"."name"
        ) FILTER (WHERE "t"."id" IS NOT NULL),
        '[]'::json
    ) AS "tags"
FROM "articles" "a"
LEFT JOIN "article_tags" "at" ON "a"."id" = "at"."article_id"
LEFT JOIN "tags" "t" ON "at"."tag_id" = "t"."id"
WHERE "a"."id" = $1
GROUP BY "a"."id";

-- name: ListArticlesWithAllRelations :many
SELECT
    "a".*,
    "c"."name" AS "category_name",
    "c"."slug" AS "category_slug",
    "sm"."description" AS "seo_description",
    "sm"."keywords" AS "seo_keywords",
    "sm"."og_image" AS "seo_og_image",
    COALESCE(
        json_agg(
            DISTINCT jsonb_build_object(
                'id', "t"."id",
                'name', "t"."name",
                'slug', "t"."slug"
            )
        ) FILTER (WHERE "t"."id" IS NOT NULL),
        '[]'::json
    ) AS "tags"
FROM "articles" "a"
LEFT JOIN "categories" "c" ON "a"."category_id" = "c"."id"
LEFT JOIN "article_seo_metadata" "sm" ON "a"."id" = "sm"."article_id"
LEFT JOIN "article_tags" "at" ON "a"."id" = "at"."article_id"
LEFT JOIN "tags" "t" ON "at"."tag_id" = "t"."id"
WHERE "a"."status" = 'published'
GROUP BY "a"."id", "c"."name", "c"."slug", "sm"."description", "sm"."keywords", "sm"."og_image"
ORDER BY "a"."published_at" DESC
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');
