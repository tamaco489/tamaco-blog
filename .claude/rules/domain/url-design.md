# URL設計・ルーティング仕様

## URL設計方針

### **設計原則**
- **SEOフレンドリー**: 検索エンジンに理解しやすいURL構造
- **ユーザーフレンドリー**: 人間が理解しやすく覚えやすいURL
- **階層構造**: 論理的なサイト構造を反映
- **永続性**: URLが変更されにくい設計
- **国際化対応**: 将来の多言語化を考慮

### **命名規則**
- **小文字使用**: すべて小文字で統一
- **ハイフン区切り**: 単語区切りはハイフン（-）を使用
- **日本語回避**: URLパラメータは英数字のみ
- **短縮形**: 適度に短く、意味の分かるURL

## フロントエンドページURL一覧

### **パブリックページ**

#### **トップ・基本ページ**
```
/                           # トップページ
/about                      # プロフィール・サイト概要
/contact                    # お問い合わせ
/privacy                    # プライバシーポリシー
/terms                      # 利用規約
/sitemap                    # サイトマップ
```

#### **記事関連ページ**
```
/articles                   # 記事一覧（最新順）
/articles/page/[page]       # 記事一覧ページネーション
/articles/[slug]            # 記事詳細ページ
/articles/popular           # 人気記事一覧
/articles/recent            # 最新記事一覧
```

#### **カテゴリ関連ページ**
```
/categories                 # カテゴリ一覧
/categories/[slug]          # カテゴリ別記事一覧
/categories/[slug]/page/[page] # カテゴリ別記事一覧ページネーション
```

#### **タグ関連ページ**
```
/tags                       # タグ一覧
/tags/[slug]                # タグ別記事一覧
/tags/[slug]/page/[page]    # タグ別記事一覧ページネーション
```

#### **アーカイブページ**
```
/archive                    # アーカイブトップ
/archive/[year]             # 年別アーカイブ
/archive/[year]/[month]     # 月別アーカイブ
/archive/[year]/page/[page] # 年別アーカイブページネーション
/archive/[year]/[month]/page/[page] # 月別アーカイブページネーション
```

#### **検索・フィード関連**
```
/search                     # 検索ページ
/search?q=[keyword]         # 検索結果
/search?q=[keyword]&page=[page] # 検索結果ページネーション
/feed.xml                   # RSS フィード
/sitemap.xml                # XMLサイトマップ
/robots.txt                 # robots.txt
```

### **管理ページ（認証必要）**

#### **ダッシュボード**
```
/admin                      # 管理画面トップ
/admin/dashboard            # ダッシュボード
/admin/analytics            # アクセス解析
```

#### **記事管理**
```
/admin/articles             # 記事一覧
/admin/articles/new         # 新規記事作成
/admin/articles/[id]/edit   # 記事編集
/admin/articles/[id]/preview # 記事プレビュー
/admin/articles/drafts      # 下書き一覧
/admin/articles/published   # 公開済み記事一覧
```

#### **カテゴリ・タグ管理**
```
/admin/categories           # カテゴリ管理
/admin/categories/new       # カテゴリ作成
/admin/categories/[id]/edit # カテゴリ編集
/admin/tags                 # タグ管理
/admin/tags/new             # タグ作成
/admin/tags/[id]/edit       # タグ編集
```

#### **メディア・設定**
```
/admin/media                # メディアライブラリ
/admin/settings             # サイト設定
/admin/settings/seo         # SEO設定
/admin/settings/profile     # プロフィール設定
```

## Next.js App Router ファイル構造

### **パブリックページ構造**
```
app/
├── page.tsx                          # / (トップページ)
├── about/
│   └── page.tsx                      # /about
├── contact/
│   └── page.tsx                      # /contact
├── articles/
│   ├── page.tsx                      # /articles
│   ├── [slug]/
│   │   └── page.tsx                  # /articles/[slug]
│   ├── page/
│   │   └── [page]/
│   │       └── page.tsx              # /articles/page/[page]
│   ├── popular/
│   │   └── page.tsx                  # /articles/popular
│   └── recent/
│       └── page.tsx                  # /articles/recent
├── categories/
│   ├── page.tsx                      # /categories
│   └── [slug]/
│       ├── page.tsx                  # /categories/[slug]
│       └── page/
│           └── [page]/
│               └── page.tsx          # /categories/[slug]/page/[page]
├── tags/
│   ├── page.tsx                      # /tags
│   └── [slug]/
│       ├── page.tsx                  # /tags/[slug]
│       └── page/
│           └── [page]/
│               └── page.tsx          # /tags/[slug]/page/[page]
├── archive/
│   ├── page.tsx                      # /archive
│   ├── [year]/
│   │   ├── page.tsx                  # /archive/[year]
│   │   ├── [month]/
│   │   │   ├── page.tsx              # /archive/[year]/[month]
│   │   │   └── page/
│   │   │       └── [page]/
│   │   │           └── page.tsx      # /archive/[year]/[month]/page/[page]
│   │   └── page/
│   │       └── [page]/
│   │           └── page.tsx          # /archive/[year]/page/[page]
├── search/
│   └── page.tsx                      # /search
├── privacy/
│   └── page.tsx                      # /privacy
├── terms/
│   └── page.tsx                      # /terms
├── sitemap/
│   └── page.tsx                      # /sitemap
├── feed.xml/
│   └── route.ts                      # /feed.xml (API Route)
├── sitemap.xml/
│   └── route.ts                      # /sitemap.xml (API Route)
└── robots.txt/
    └── route.ts                      # /robots.txt (API Route)
```

### **管理ページ構造**
```
app/admin/
├── page.tsx                          # /admin (リダイレクト)
├── dashboard/
│   └── page.tsx                      # /admin/dashboard
├── analytics/
│   └── page.tsx                      # /admin/analytics
├── articles/
│   ├── page.tsx                      # /admin/articles
│   ├── new/
│   │   └── page.tsx                  # /admin/articles/new
│   ├── [id]/
│   │   ├── edit/
│   │   │   └── page.tsx              # /admin/articles/[id]/edit
│   │   └── preview/
│   │       └── page.tsx              # /admin/articles/[id]/preview
│   ├── drafts/
│   │   └── page.tsx                  # /admin/articles/drafts
│   └── published/
│       └── page.tsx                  # /admin/articles/published
├── categories/
│   ├── page.tsx                      # /admin/categories
│   ├── new/
│   │   └── page.tsx                  # /admin/categories/new
│   └── [id]/
│       └── edit/
│           └── page.tsx              # /admin/categories/[id]/edit
├── tags/
│   ├── page.tsx                      # /admin/tags
│   ├── new/
│   │   └── page.tsx                  # /admin/tags/new
│   └── [id]/
│       └── edit/
│           └── page.tsx              # /admin/tags/[id]/edit
├── media/
│   └── page.tsx                      # /admin/media
└── settings/
    ├── page.tsx                      # /admin/settings
    ├── seo/
    │   └── page.tsx                  # /admin/settings/seo
    └── profile/
        └── page.tsx                  # /admin/settings/profile
```

## URL パラメータ仕様

### **動的ルートパラメータ**
```typescript
// URL パラメータの型定義
interface PageParams {
  // 記事詳細
  slug: string;                       // 記事スラッグ

  // ページネーション
  page: string;                       // ページ番号（数値文字列）

  // アーカイブ
  year: string;                       // 年（YYYY）
  month: string;                      // 月（MM）

  // 管理画面
  id: string;                         // UUID
}

// クエリパラメータの型定義
interface SearchParams {
  q?: string;                         // 検索キーワード
  category?: string;                  // カテゴリフィルタ
  tag?: string;                       // タグフィルタ
  sort?: 'latest' | 'popular' | 'oldest'; // ソート順
  year?: string;                      // 年フィルタ
  month?: string;                     // 月フィルタ
  page?: string;                      // ページ番号
}
```

### **URL 生成ヘルパー関数**
```typescript
// lib/urls.ts
export const urls = {
  // パブリックページ
  home: () => '/',
  about: () => '/about',
  contact: () => '/contact',

  // 記事関連
  articles: (page?: number) => page ? `/articles/page/${page}` : '/articles',
  article: (slug: string) => `/articles/${slug}`,
  articlesByCategory: (categorySlug: string, page?: number) =>
    page ? `/categories/${categorySlug}/page/${page}` : `/categories/${categorySlug}`,
  articlesByTag: (tagSlug: string, page?: number) =>
    page ? `/tags/${tagSlug}/page/${page}` : `/tags/${tagSlug}`,

  // アーカイブ
  archive: () => '/archive',
  archiveByYear: (year: number, page?: number) =>
    page ? `/archive/${year}/page/${page}` : `/archive/${year}`,
  archiveByMonth: (year: number, month: number, page?: number) =>
    page ? `/archive/${year}/${month.toString().padStart(2, '0')}/page/${page}` :
           `/archive/${year}/${month.toString().padStart(2, '0')}`,

  // 検索
  search: (query?: string, page?: number) => {
    const params = new URLSearchParams();
    if (query) params.set('q', query);
    if (page && page > 1) params.set('page', page.toString());
    return `/search${params.toString() ? `?${params.toString()}` : ''}`;
  },

  // 管理画面
  admin: {
    dashboard: () => '/admin/dashboard',
    articles: () => '/admin/articles',
    articleNew: () => '/admin/articles/new',
    articleEdit: (id: string) => `/admin/articles/${id}/edit`,
    categories: () => '/admin/categories',
    categoryEdit: (id: string) => `/admin/categories/${id}/edit`,
    tags: () => '/admin/tags',
    tagEdit: (id: string) => `/admin/tags/${id}/edit`,
    media: () => '/admin/media',
    settings: () => '/admin/settings'
  }
} as const;
```

## SEO 考慮事項

### **URL 最適化**
- **スラッグ生成**: 日本語タイトルから英数字スラッグを自動生成
- **パンくずリスト**: 階層構造をスキーママークアップで表現
- **カノニカルURL**: 重複コンテンツ回避のため適切なcanonical設定
- **OGP URL**: SNSシェア時の正確なURL設定

### **リダイレクト設定**
```typescript
// next.config.mjs
const redirects = [
  // レガシーURL対応
  {
    source: '/post/:slug',
    destination: '/articles/:slug',
    permanent: true
  },
  // 管理画面ルート
  {
    source: '/admin',
    destination: '/admin/dashboard',
    permanent: false
  }
];
```
