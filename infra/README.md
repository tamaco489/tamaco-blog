# Infrastructure

このディレクトリはTerraformを使用してAWSリソースを管理します。

## 概要

AWS上でブログシステムのインフラを構築・管理するためのTerraformモジュール群です。
各サービス毎にディレクトリが分かれており、独立してデプロイ可能な構成になっています。

## 基本操作

各ディレクトリで以下のコマンドを実行してください。

### 初期化
```bash
make init ENV=dev
```

### 実行計画確認
```bash
make plan ENV=dev
```

### 変更適用
```bash
make apply ENV=dev
```

## コード品質管理

### フォーマット
```bash
make fmt
```

### リント
```bash
make lint
```

### セキュリティスキャン
```bash
make security
```

## 環境

- `ENV=dev` - 開発環境
- `ENV=prd` - 本番環境