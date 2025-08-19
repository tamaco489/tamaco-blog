# CLAUDE.md

このファイルは Claude Code (claude.ai/code) がこのリポジトリで作業する際のガイダンスを提供します。

@.claude/rules/dev-rules/globals.md
@.claude/rules/dev-rules/techstack.md
@.claude/rules/dev-rules/frontend-structure.md
@.claude/rules/dev-rules/openapi-spec.md
@.claude/rules/domain/blog-domain.md
@.claude/rules/domain/features.md
@.claude/rules/domain/data-models.md
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
make install-tools      # 環境変数ファイルの作成、開発ツールのインストール
make up                 # Docker環境起動（PostgreSQL含む）
make migrate-up         # データベースマイグレーション
make logs               # APIサーバのログ出力（起動確認）
```

#### OpenAPI 開発

```bash
cd backend/api/article
make bundle-openapi      # OpenAPI specのバンドル
make gen-api            # APIインターフェースと型定義の生成
```

## アーキテクチャ概要

- **フロントエンド**: Next.js 14 App Router + TypeScript + TailwindCSS
- **バックエンド**: Go + AWS Lambda（予定）
- **データベース**: Supabase PostgreSQL
- **ホスティング**: AWS S3 + CloudFront
- **CI/CD**: GitHub Actions
