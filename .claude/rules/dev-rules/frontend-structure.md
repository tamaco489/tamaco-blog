# フロントエンド ディレクトリ構成と開発ルール

## 現行ディレクトリ構成

以下は `frontend` 配下の詳細な構成です。実際の構成と差異が生じた場合は、このドキュメントを更新してください。

```text
frontend/
├─ app/
│  ├─ favicon.ico
│  ├─ globals.css
│  ├─ layout.tsx
│  └─ page.tsx
├─ components/
├─ features/
├─ hooks/
├─ providers/
├─ utils/
├─ constants/
├─ types/
├─ styles/
├─ lib/
├─ tests/
├─ next-env.d.ts
├─ next.config.mjs
├─ postcss.config.mjs
├─ tailwind.config.ts
├─ tsconfig.json
├─ package.json
├─ package-lock.json
├─ README.md
└─ node_modules/
```

## 各ディレクトリの詳細な役割

### `app/` - Next.js App Router
- **app/layout.tsx**: 全ページ共通のレイアウト（メタ情報やレイアウト枠組み）
- **app/page.tsx**: ルート (`/`) のページコンポーネント
- **app/globals.css**: グローバルスタイル。TailwindCSS のベースレイヤ等を読み込み
- 各ルート配下に `page.tsx`、`layout.tsx`、必要に応じて `loading.tsx` などを配置
- **ルール**: ビジネスロジックは配置しない、ルーティング専用

### `components/` - 横断的UIコンポーネント
- ドメインに依存しない再利用可能なUIコンポーネントを配置
- Button, Modal, Layout などの汎用コンポーネント
- 例: `components/ui/Button.tsx`, `components/layout/Header.tsx`

### `features/` - 機能別集約ディレクトリ
- 機能（ドメイン）単位でUI・hooks・ユーティリティ・型などを集約
- 例: `features/blog/`, `features/auth/`, `features/user/`
- 各featureディレクトリ内に `components/`, `hooks/`, `utils/`, `types/` を配置可能

### `hooks/` - 横断的カスタムフック
- ドメインに依存しない横断的に使うカスタムフックを配置
- 例: `useLocalStorage`, `useDebounce`, `useApi`

### `providers/` - Reactコンテキスト/プロバイダー
- ReactのContextやアプリケーション全体で利用するプロバイダーを配置
- 例: `ThemeProvider`, `AuthProvider`, `QueryProvider`

### `utils/` - 汎用関数
- ドメインに依存しない汎用関数を配置
- 例: `formatDate`, `validateEmail`, `debounce`

### `constants/` - グローバル定数
- グローバルに利用する定数を配置
- 例: API endpoints, 設定値, 列挙値

### `types/` - 横断的型定義
- 横断的な型定義を配置
- 例: `api.ts`, `common.ts`, `env.ts`

### `styles/` - スタイル関連
- スタイル関連（CSS/Tailwindの設定拡張など）を配置
- カスタムCSSファイルやTailwind拡張

### `lib/` - ライブラリラッパー
- ライブラリ/標準処理の共通化コードや外部連携の薄いラッパー等を配置
- 例: `api.ts`, `auth.ts`, `database.ts`

### `tests/` - 自動テスト
- 自動テスト（ユニット/統合）関連を配置
- `__tests__/`, `__mocks__/` などのサブディレクトリ

## 設定ファイル

- **next.config.mjs**: Next.js のビルド/実行設定
- **postcss.config.mjs**: PostCSS/TailwindCSS 設定
- **tailwind.config.ts**: TailwindCSS のプロジェクト設定
- **tsconfig.json**: TypeScript 設定
- **next-env.d.ts**: Next.js により自動生成される型定義のエントリ

## 開発ルール

### ファイル命名規則
- **kebab-case**: ファイル名とディレクトリ名
- **PascalCase**: コンポーネント名
- **camelCase**: 関数名、変数名

### コンポーネント設計原則
- **関数コンポーネント優先**
- **TypeScript interface使用**
- **default export使用**
- **Server Components優先**

### スタイリング方針
- **TailwindCSS優先**: メインのスタイリング手法
- **カスタムCSS最小限**: 必要最小限に抑制

## 変更時のガイドライン

- **ドキュメント更新**: ディレクトリ／ファイルを追加・削除した場合は、本ドキュメントのツリーを更新
- **事前承認**: UI/UX の変更が絡む場合は、プロジェクトのルールに従い事前に承認を得る
- **ライブラリ管理**: ライブラリの追加やバージョン変更は、目的・影響範囲（ビルドサイズ/セキュリティ/保守性）を明記のうえ相談・承認を得る

## 技術前提

- **フレームワーク**: Next.js 14（App Router）
- **言語**: TypeScript / React 18
- **スタイリング**: TailwindCSS 3

