# API 設計（将来のバックエンド用: 後に Open API Spec で定義予定）

### **RESTful API エンドポイント**

#### **記事関連 API**

```
GET    /api/articles              # 記事一覧取得
GET    /api/articles/:slug        # 記事詳細取得
POST   /api/articles              # 記事作成（認証必要）
PUT    /api/articles/:id          # 記事更新（認証必要）
DELETE /api/articles/:id          # 記事削除（認証必要）
PATCH  /api/articles/:id/publish  # 記事公開（認証必要）
PATCH  /api/articles/:id/views    # 閲覧数更新
```

#### **カテゴリ関連 API**

```
GET    /api/categories            # カテゴリ一覧取得
GET    /api/categories/:slug      # カテゴリ詳細取得
GET    /api/categories/:slug/articles # カテゴリ別記事一覧
POST   /api/categories            # カテゴリ作成（認証必要）
PUT    /api/categories/:id        # カテゴリ更新（認証必要）
DELETE /api/categories/:id        # カテゴリ削除（認証必要）
```

#### **タグ関連 API**

```
GET    /api/tags                  # タグ一覧取得
GET    /api/tags/:slug            # タグ詳細取得
GET    /api/tags/:slug/articles   # タグ別記事一覧
POST   /api/tags                  # タグ作成（認証必要）
PUT    /api/tags/:id              # タグ更新（認証必要）
DELETE /api/tags/:id              # タグ削除（認証必要）
```

#### **検索・フィルタ API**

```
GET    /api/search?q=keyword      # 記事検索
GET    /api/articles/archive/:year/:month # 月別アーカイブ
GET    /api/articles/popular      # 人気記事
GET    /api/articles/recent       # 最新記事
```

### **クエリパラメータ仕様**

```typescript
// 一覧取得共通パラメータ
interface ListParams {
  page?: number; // ページ番号（デフォルト: 1）
  limit?: number; // 取得件数（デフォルト: 10）
  sort?: string; // ソート項目（created_at, updated_at, view_count）
  order?: "asc" | "desc"; // ソート順（デフォルト: desc）
}

// 記事一覧固有パラメータ
interface ArticleListParams extends ListParams {
  status?: "published" | "draft" | "private";
  category?: string; // カテゴリスラッグ
  tag?: string; // タグスラッグ
  search?: string; // 検索キーワード
  year?: number; // 年指定
  month?: number; // 月指定
}
```
