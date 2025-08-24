# インフラ ディレクトリ構成ルール

## 共通ディレクトリ構成

各 AWS リソースディレクトリは以下の共通構成に従います：

```
{service}/
├── Makefile              # Terraform操作コマンド
├── provider.tf           # AWSプロバイダー設定
├── variables.tf          # 変数定義
├── {main}.tf             # メインリソース定義
├── outputs.tf            # 出力値定義（必要に応じて）
├── data.tf               # データソース定義（必要に応じて）
├── tfbackend/            # Terraform Backend設定
│   ├── dev.tfbackend     # 開発環境用
│   └── prd.tfbackend     # 本番環境用
└── tfvars/               # 環境変数ファイル
    ├── dev.tfvars        # 開発環境変数
    └── prd.tfvars        # 本番環境変数
```

## ファイル役割

### 必須ファイル

- **`Makefile`**: 各ディレクトリで共通設定を参照
  ```makefile
  include ../make/base.mk
  BACKEND_CONFIG := -backend-config="./tfbackend/$(ENV).tfbackend"
  ```
- **`provider.tf`**: AWS Provider 設定、リージョン指定
- **`variables.tf`**: 環境固有の変数定義（environment, tags 等）
- **`{main}.tf`**: 当該サービスのメインリソース定義

### Backend・環境設定

- **`tfbackend/`**: S3 バックエンド設定（状態ファイル保存先）
- **`tfvars/`**: 環境別の具体的な変数値

### オプションファイル

- **`outputs.tf`**: 他モジュールで参照する値の出力
- **`data.tf`**: 他の Terraform で作成済みリソースの参照

## Makefile 構成

### 共通設定（`infra/make/base.mk`）
- 基本的な Terraform コマンド（init, plan, apply, destroy）
- AWS Profile・Region 設定
- フォーマット・検証・セキュリティチェック
- 暗号化・復号化コマンド（sops）

### 各ディレクトリ（例：`infra/acm/Makefile`）
```makefile
include ../make/base.mk
BACKEND_CONFIG := -backend-config="./tfbackend/$(ENV).tfbackend"
```

基本的に `base.mk` をインクルードし、モジュール固有の `BACKEND_CONFIG` のみを設定する
