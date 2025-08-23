# バックエンド ディレクトリ構成と開発ルール

## 現行ディレクトリ構成

以下は `backend` 配下の詳細な構成です。実際の構成と差異が生じた場合は、このドキュメントを更新してください。

```text
backend/
└── api/
    ├── compose.yaml          # Docker Compose 設定
    └── article/              # 記事管理 API サービス
        ├── Makefile          # 開発・ビルド・テスト用コマンド
        ├── go.mod            # Go モジュール定義
        ├── go.sum            # 依存関係のハッシュ値
        ├── cmd/              # アプリケーションエントリーポイント
        │   └── main.go       # メイン関数
        ├── internal/         # 内部パッケージ（外部からアクセス不可）
        │   ├── controller/   # HTTP ハンドラー層
        │   ├── domain/       # ドメインロジック
        │   ├── gen/          # OpenAPI 自動生成ファイル
        │   ├── handler/      # ルーティング設定
        │   ├── library/      # 共通ライブラリ
        │   ├── repository/   # データアクセス層
        │   └── usecase/      # ユースケース層
        ├── spec/             # OpenAPI 仕様定義
        │   ├── openapi.yaml  # 自動生成バンドル仕様（編集禁止）
        │   ├── openapi_base.yaml # ベース仕様定義
        │   ├── api/          # 仕様分割ファイル群
        │   └── templates/    # コード生成テンプレート
        ├── build/            # ビルド成果物
        ├── scripts/          # ユーティリティスクリプト
        └── tmp/              # 一時ファイル
```

## 各ディレクトリの詳細な役割

### `api/` - API サービス群

- マイクロサービス形式で各 API を分離配置
- 各 API サービスは独立したモジュールとして管理

### `api/article/` - 記事管理 API サービス

- 記事の CRUD 操作を提供する Go アプリケーション
- クリーンアーキテクチャに基づく設計

#### `cmd/` - エントリーポイント

- **main.go**: アプリケーションのエントリーポイント
- 依存関係の注入とサーバー起動処理

#### `internal/` - 内部パッケージ（Clean Architecture）

- Go の internal パッケージ規約に従い、外部からのアクセスを制限

#### `internal/controller/` - プレゼンテーション層

- HTTP ハンドラー、ルーティング
- リクエスト/レスポンスの変換
- 入力値検証

#### `internal/domain/` - ドメイン層

- ビジネスロジック、エンティティ
- ドメインサービス
- インターフェース定義

#### `internal/gen/` - 自動生成コード

- `server.gen.go`: OpenAPI からの自動生成サーバーインターフェース
- `types.gen.go`: OpenAPI からの自動生成型定義
- **重要**: **手動編集厳禁** - `make gen-api` で再生成される

#### `internal/handler/` - ルーティング設定

- HTTP ルーティングの設定
- ミドルウェアの適用
- サーバー起動処理

#### `internal/library/` - 共通ライブラリ

- `config/`: 環境変数・設定管理、AWS Secrets Manager 連携
- `logger/`: 構造化ログ出力、コンテキスト管理

#### `internal/repository/` - データアクセス層

- データベースアクセス実装
- 外部 API 連携
- キャッシュ処理

#### `internal/usecase/` - アプリケーション層

- ビジネスユースケースの実装
- ドメイン層と repository 層の調整
- トランザクション管理

#### `build/` - ビルド成果物

- コンパイル済みバイナリ
- Docker イメージビルド用

#### `scripts/` - 運用スクリプト

- デプロイメントスクリプト
- 開発環境セットアップ

#### `scripts/localstack/` - ローカル開発環境

- LocalStack（AWS サービスローカル実行）設定
- 開発用データベース初期化

#### `spec/` - OpenAPI 仕様定義

- `openapi_base.yaml`: ベース仕様定義（手動編集対象）
- `openapi.yaml`: 自動生成バンドル仕様（**編集禁止**）
- `api/paths/`: エンドポイント定義（articles.yaml, categories.yaml, tags.yaml, search.yaml）
- `api/schemas/`: データモデル定義（Article.yaml, Category.yaml, Tag.yaml）
- `api/requests/`: リクエストボディ定義
- `api/responses/`: レスポンス定義
- `api/parameters/`: 共通パラメータ定義
- `templates/`: コード生成テンプレート

## 開発ルール

### アーキテクチャ設計原則

- **Clean Architecture**: レイヤー分離による保守性向上
- **DDD**: ドメイン駆動設計によるビジネスロジック集約
- **SOLID 原則**: 拡張性・保守性を重視した設計

### ファイル命名規則

- **snake_case**: ファイル名、ディレクトリ名
- **PascalCase**: 構造体名、インターフェース名
- **camelCase**: 関数名、変数名、定数

### パッケージ設計原則

- **internal パッケージ**: 外部アクセス制限の活用
- **インターフェース分離**: 依存関係の逆転
- **単一責任**: 各パッケージは明確な責任を持つ

### コード品質管理

- **Go 標準**: gofmt、golint、go vet の活用
- **テストカバレッジ**: 80% 以上を目標
- **エラーハンドリング**: 適切なエラー処理の実装

### セキュリティ対策

- **入力検証**: すべての入力値の検証
- **認証・認可**: JWT トークンベース認証
- **ログ管理**: セキュリティイベントのログ記録

## 変更時のガイドライン

- **ドキュメント更新**: ディレクトリ／ファイルを追加・削除した場合は、本ドキュメントのツリーを更新
- **事前承認**: アーキテクチャ変更は設計レビューと承認が必要
- **依存関係管理**: ライブラリの追加やバージョン変更は、セキュリティ・パフォーマンス影響を評価

## 基本コマンド

### 開発環境セットアップ
```bash
cd backend/api/article
make install-tools  # 環境変数ファイルの作成、開発ツールのインストール
make up             # Docker環境起動（PostgreSQL含む）
make migrate-up     # データベースマイグレーション
make logs          # APIサーバのログ出力（起動確認）
```

### 開発・ビルド
```bash
make build          # ビルド実行
make run           # ローカル実行
make test          # テスト実行
make lint          # リント実行
```

### OpenAPI 関連
```bash
make bundle-openapi # OpenAPI仕様のバンドル
make gen-api       # APIインターフェースと型定義の生成
make serve-docs    # API ドキュメント表示
```

## 技術前提

- **言語**: Go 1.25+
- **フレームワーク**: 標準ライブラリベース（oapi-codegen による生成）
- **データベース**: PostgreSQL（Supabase）
- **クラウド**: AWS Lambda + API Gateway（予定）
- **コンテナ**: Docker
- **CI/CD**: GitHub Actions

