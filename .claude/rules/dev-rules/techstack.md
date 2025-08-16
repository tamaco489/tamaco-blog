# 技術スタック定義

## 開発言語・実装ツール
- **Go**: サーバーサイド言語
- **TypeScript**: フロントエンド言語
- **Terraform**: Infrastructure as Code

## Webフレームワーク
- **Next.js 14**: App Router使用（SSG + ISR）

## UI/スタイリング
- **TailwindCSS**: メインのスタイリングフレームワーク
- **カスタムCSS**: 最小限に抑制

## データストア
- **Supabase**: PostgreSQL データベース
- **DynamoDB**: 将来実装予定（初期段階では未実装）

## クラウドサービス
- **AWS S3**: 静的ファイルホスティング
- **AWS CloudFront**: CDN配信
- **AWS API Gateway**: APIルーティング
- **AWS Lambda**: サーバーレス関数
- **AWS ECR**: コンテナリポジトリ

## CI/CD
- **GitHub Actions**: 自動ビルド・デプロイ

## 監視・分析
- **New Relic**: アプリケーション監視
- **Google Analytics**: ユーザー行動分析

## バージョン管理ルール
- **勝手なバージョン変更禁止**: 理由と承認が必要
- **セキュリティアップデートは例外**: 即座に適用可能
