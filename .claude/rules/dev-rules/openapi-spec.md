# OpenAPI Specification ルール

## 概要

このドキュメントは、tamaco-blog バックエンド API の OpenAPI Specification（OAS）の設計・実装に関するルールと構造を定義します。

## OpenAPI バージョン

- **バージョン**: OpenAPI 3.0.0
- **理由**: 広くサポートされており、コード生成ツールとの互換性が高い

## ディレクトリ構造

```
backend/api/{service-name}/spec/
├── openapi.yaml              # 自動生成ファイル（手動編集禁止）
├── openapi_base.yaml         # ベースとなるAPI定義（手動編集対象）
└── api/
    ├── paths/                # エンドポイント定義
    │   ├── articles.yaml     # /articles エンドポイント群
    │   ├── categories.yaml   # /categories エンドポイント群
    │   └── tags.yaml         # /tags エンドポイント群
    ├── parameters/           # 共通パラメータ定義
    │   ├── common.yaml       # 汎用パラメータ
    │   └── pagination.yaml   # ページネーション用
    ├── requests/             # リクエストボディ定義
    │   ├── ArticleCreate.yaml
    │   ├── ArticleUpdate.yaml
    │   └── ...
    ├── responses/            # レスポンス定義
    │   ├── HealthCheck.yaml
    │   ├── BadRequest.yaml
    │   ├── Unauthorized.yaml
    │   ├── Forbidden.yaml
    │   ├── NotFound.yaml
    │   ├── AlreadyExists.yaml
    │   ├── InternalServerError.yaml
    │   └── ...
    └── schemas/              # データモデル定義
        ├── Article.yaml
        ├── Category.yaml
        ├── Tag.yaml
        └── ...
```

## ファイル分割ルール

### 1. **openapi_base.yaml**

- メインの API 定義ファイル
- `info`、`servers`、`security`、基本的な`paths`を含む
- 他の YAML ファイルを`$ref`で参照

### 2. **paths/**

- エンドポイント毎にファイルを分割
- リソース単位でグループ化
- 例: `/articles`関連は`articles.yaml`に集約

### 3. **parameters/**

- 共通で使用されるパラメータを定義
- ページネーション、フィルタリング、ソート等

### 4. **requests/**

- リクエストボディのスキーマを定義
- 操作とリソースを組み合わせた命名（例: `ArticleCreate.yaml`）

### 5. **responses/**

- レスポンススキーマを定義
- HTTP ステータスコードに対応
- エラーレスポンスは共通化

### 6. **schemas/**

- データモデルの定義
- エンティティ単位でファイル作成

## 命名規則

### スキーマ名

- **PascalCase**を使用
- 意味が明確な名前を付ける
- 例: `Article`, `CategoryList`, `ErrorResponse`

### operationId

- **camelCase**を使用
- 動詞 + リソース名の形式
- 例: `getArticles`, `createArticle`, `updateArticleById`

### パラメータ名

- **snake_case**を使用（URL パラメータ、クエリパラメータ）
- 例: `article_id`, `page_size`, `sort_order`

### プロパティ名（JSON）

- **snake_case**を使用
- 例: `article_id`, `created_at`, `updated_at`

## 記述ルール

### 1. **必須項目**

各定義には以下を必ず含める：

- `title`: スキーマのタイトル
- `description`: 日本語での詳細説明
- `type`: データ型
- `example`: 実例（理解しやすいデータ）

### 2. **エラーレスポンス**

統一されたエラーレスポンス形式：

```yaml
type: object
required:
  - code
  - message
properties:
  code:
    type: string
    description: エラーコード
  message:
    type: string
    description: エラーメッセージ
  details:
    type: array
    description: 詳細情報（オプション）
```

### 3. **ページネーション**

一覧取得エンドポイントには必ずページネーションを実装：

```yaml
parameters:
  - name: page
    in: query
    type: integer
    default: 1
  - name: page_size
    in: query
    type: integer
    default: 20
    maximum: 100
```

### 4. **日時フォーマット**

- ISO 8601 形式を使用
- 例: `2024-01-01T09:00:00Z`
- フィールド名: `created_at`, `updated_at`, `published_at`

### 5. **認証・認可**

- Bearer Token (JWT) を使用
- ヘルスチェック以外のエンドポイントは原則認証必須
- 公開 API は`security: []`で明示的に認証不要を指定

## バリデーションルール

### 文字列

```yaml
title:
  type: string
  minLength: 1
  maxLength: 200
  pattern: '^[^\s].*[^\s]$' # 前後の空白を禁止
```

### 数値

```yaml
page_size:
  type: integer
  minimum: 1
  maximum: 100
  default: 20
```

### 配列

```yaml
tags:
  type: array
  minItems: 0
  maxItems: 10
  uniqueItems: true
```

## コード生成

### 生成コマンド

```bash
# OpenAPI spec のバンドル（開発時）
make bundle-openapi

# APIサーバーインターフェースと型定義の生成
make gen-api
```

### 生成ファイル

- `internal/gen/server.gen.go`: サーバー側インターフェース
- `internal/gen/types.gen.go`: 型定義
- **これらは自動生成ファイルのため手動編集禁止**

## ベストプラクティス

1. **段階的な実装**

   - まずスキーマを定義
   - 次にエンドポイントを定義
   - 最後に実装

2. **再利用性**

   - 共通コンポーネントは`components`に定義
   - `$ref`で参照して重複を避ける

3. **バージョニング**

   - URL パスにバージョン含める（`/v1/`）
   - 破壊的変更は新バージョンで対応

4. **ドキュメント**

   - すべての要素に`description`を記載
   - `example`で具体例を提供
   - 日本語で分かりやすく記述

5. **テスト**
   - OpenAPI Spec のバリデーション実施
   - モックサーバーでの動作確認
   - クライアント生成コードでのテスト

## 変更時の手順

1. **仕様変更の場合**

   ```bash
   # 1. openapi_base.yaml または api/ 配下のファイルを編集

   # 2. スペックのバンドル
   make bundle-openapi

   # 3. コード生成
   make gen-api
   ```

2. **新規エンドポイント追加**

   - 適切な paths ファイルに追加
   - 必要なスキーマを定義
   - バンドル・生成・実装

3. **スキーマ変更**
   - 後方互換性を考慮
   - 必須フィールドの追加は避ける
   - 非推奨化は`deprecated: true`を使用

## 参考リンク

- [OpenAPI Specification 3.0](https://swagger.io/specification/v3/)
- [oapi-codegen](https://github.com/deepmap/oapi-codegen)
- [OpenAPI Generator](https://openapi-generator.tech/)
