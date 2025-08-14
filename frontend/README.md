### tamaco-blog frontend

## フロントエンド ディレクトリ構成

以下は `frontend` 配下のディレクトリ構成です。

```text
frontend/
├─ app/
├─ components/
├─ features/
├─ hooks/
├─ providers/
├─ utils/
├─ constants/
├─ types/
├─ styles/
├─ lib/
└─ tests/
```

### 各ディレクトリの役割

- `app/`: Next.js App Router のアプリケーションルート。Routing files のみを配置します。
- `components/`: 横断的（ドメイン非依存）の再利用可能な UI コンポーネント。
- `features/`: 機能（ドメイン）単位で UI・hooks・ユーティリティ・型などを集約。
- `hooks/`: ドメインに依存しない横断的に使うカスタムフック。
- `providers/`: アプリケーション全体で利用する React のコンテキスト/プロバイダー群。
- `utils/`: ドメインに依存しない汎用関数。
- `constants/`: グローバルに利用する定数。
- `types/`: 横断的な型定義。
- `styles/`: スタイル関連（CSS/Tailwind の設定拡張など）。
- `lib/`: ライブラリ/標準処理の共通化コードや外部連携の薄いラッパー等。
- `tests/`: 自動テスト（ユニット/統合）。

## 起動方法

### セットアップ

```bash
cd frontend
npm ci
```

### 開発サーバー起動

```bash
npm run dev
```

### 本番ビルド・起動

```bash
npm run build
npm run start
```

### Lint 実行

```bash
npm run lint
```
