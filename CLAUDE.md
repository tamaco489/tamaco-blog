# CLAUDE.md

このファイルは Claude Code (claude.ai/code) がこのリポジトリで作業する際のガイダンスを提供します。

@.claude/rules/dev-rules/globals.md
@.claude/rules/dev-rules/techstack.md
@.claude/rules/dev-rules/frontend-structure.md
@.claude/rules/dev-rules/backend-structure.md
@.claude/rules/dev-rules/go-coding-rules.md
@.claude/rules/dev-rules/openapi-spec.md
@.claude/rules/dev-rules/file-formatting-rules.md
@.claude/rules/domain/blog-domain.md
@.claude/rules/domain/features.md
@.claude/rules/domain/url-design.md

## プロジェクト概要

Next.js を使用したブログシステムです。AWS S3 と CloudFront でホスティング、Supabase でデータ管理、Strapi で CMS 機能を提供します。モノレポ構造で frontend、backend、infrastructure コンポーネントを分離しています。

## 基本コマンド

### フロントエンド開発

フロントエンドコマンドは `frontend` ディレクトリから実行してください。

```bash
cd frontend
npm ci      # 依存関係のインストール
npm run dev # 開発サーバー起動
```

### バックエンド開発

バックエンドコマンドは `backend/api/article` ディレクトリから実行してください。

#### 初回セットアップと起動

```bash
cd backend/api/article
make install-tools # 環境変数ファイルの作成、開発ツールのインストール
make up            # Docker環境起動（PostgreSQL含む）
make migrate-up    # データベースマイグレーション
make load-masters  # マスタデータ投入
make logs          # APIサーバのログ出力（起動確認）
```

#### OpenAPI 開発

```bash
cd backend/api/article
make bundle-openapi # OpenAPI specのバンドル
make gen-api        # APIインターフェースと型定義の生成
```

#### データベース開発手順

##### 1. マイグレーションファイルの追加

```bash
# 新しいマイグレーションファイルを作成
make migrate-create NAME=create_users_table

# マイグレーションの適用・確認
make migrate-up     # マイグレーション適用
make migrate-status # ステータス確認
make migrate-down   # ロールバック（必要時）
```

##### 2. sqlc用クエリファイルの追加

```bash
# クエリファイル作成 (internal/repository/queries/テーブル名.sql)
# sqlcコード生成
make gen-sqlc
```

##### 3. マスタデータの投入

```bash
# マスタデータ投入
make load-masters   # マスタデータ投入
make reset-masters  # 全削除してマスタデータ再投入

# 新しいマスタファイル作成 (scripts/masters/番号_テーブル名.sql)
```

## アーキテクチャ概要

- **フロントエンド**: Next.js 14 App Router + TypeScript + TailwindCSS
- **バックエンド**: Go + AWS Lambda（予定）
- **データベース**: Supabase PostgreSQL
- **ホスティング**: AWS S3 + CloudFront
- **CI/CD**: GitHub Actions
