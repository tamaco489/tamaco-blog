# マスタデータ投入スクリプト

## 使用方法

### ローカル環境（Docker Compose）

```bash
DB_HOST=localhost DB_PORT=5432 DB_NAME=core DB_USER=core DB_PASSWORD=password#0 make load-masters
```

### リモート環境（Supabase等）

```bash
DB_HOST=your-host.com DB_PORT=5432 DB_NAME=postgres DB_USER=your-user DB_PASSWORD=your-password make load-masters
```

## その他のコマンド

- `make clear-masters` - マスタデータを削除
- `make init-masters` - マスタデータをリセット＆再投入