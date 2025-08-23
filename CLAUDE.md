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

Next.js を使用したブログシステムです。AWS S3 と CloudFront でホスティング、Supabase でデータ管理、Strapi で CMS 機能を提供します。

モノレポ構造で frontend、backend、infrastructure コンポーネントを分離しています。

## アーキテクチャ概要

- **フロントエンド**: Next.js 14 App Router + TypeScript + TailwindCSS
- **バックエンド**: Go + AWS Lambda（予定）
- **データベース**: Supabase PostgreSQL
- **ホスティング**: AWS S3 + CloudFront
- **CI/CD**: GitHub Actions
