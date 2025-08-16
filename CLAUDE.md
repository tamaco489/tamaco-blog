# CLAUDE.md

このファイルは Claude Code (claude.ai/code) がこのリポジトリで作業する際のガイダンスを提供します。

@.claude/rules/dev-rules/globals.md
@.claude/rules/dev-rules/techstack.md
@.claude/rules/dev-rules/frontend-structure.md

## プロジェクト概要

Next.js を使用したブログシステムです。AWS S3 と CloudFront でホスティング、Supabase でデータ管理、Strapi で CMS 機能を提供します。モノレポ構造で frontend、backend、infrastructure コンポーネントを分離しています。

## 基本コマンド

特記がない限り、すべてのコマンドは `frontend` ディレクトリから実行してください。

### 開発ワークフロー
```bash
cd frontend
npm ci                    # 依存関係のインストール
npm run dev              # 開発サーバー起動
npm run build            # 本番ビルド
npm run start            # 本番サーバー起動
npm run lint             # リント実行
```

### 品質保証
変更後は必ず `npm run lint` を実行してコード品質と一貫性を確保してください。

## アーキテクチャ概要

- **フロントエンド**: Next.js 14 App Router + TypeScript + TailwindCSS
- **バックエンド**: Go + AWS Lambda（予定）
- **データベース**: Supabase PostgreSQL
- **ホスティング**: AWS S3 + CloudFront
- **CI/CD**: GitHub Actions
