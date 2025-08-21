-- Categories master data
INSERT INTO "categories" ("name", "slug", "description", "display_order") VALUES
('バックエンド開発', 'backend-development', 'サーバーサイド開発に関する技術記事', 1),
('フロントエンド開発', 'frontend-development', 'クライアントサイド開発に関する技術記事', 2),
('DevOps', 'devops', '開発運用・インフラに関する記事', 3),
('データベース', 'database', 'データベース設計・運用に関する記事', 4),
('プログラミング言語', 'programming-languages', 'プログラミング言語の学習・Tips', 5),
('ツール・環境構築', 'tools-setup', '開発ツールや環境構築に関する記事', 6),
('設計・アーキテクチャ', 'design-architecture', 'システム設計・アーキテクチャに関する記事', 7),
('学習記録', 'learning-notes', '技術学習の記録・メモ', 8),
('チュートリアル', 'tutorials', 'ハンズオン形式のチュートリアル記事', 9),
('その他', 'others', 'その他の技術関連記事', 99)
ON CONFLICT ("slug") DO UPDATE SET
    "name" = EXCLUDED."name",
    "description" = EXCLUDED."description",
    "display_order" = EXCLUDED."display_order",
    "updated_at" = CURRENT_TIMESTAMP;
