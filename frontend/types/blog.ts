// ブログ関連の型定義

export interface Article {
  id: string;
  title: string;
  slug: string;
  content: string;
  excerpt?: string;
  status: "draft" | "published" | "private";
  published_at?: string;
  created_at: string;
  updated_at: string;
  category_id?: string;
  author_id: string;
  view_count: number;
  meta_description?: string;
  meta_keywords?: string;
  featured_image?: string;

  // Relations
  category?: Category;
  author?: User;
  tags?: Tag[];
}

export interface Category {
  id: string;
  name: string;
  slug: string;
  description?: string;
  color?: string;
  display_order: number;
  created_at: string;
  updated_at: string;
}

export interface Tag {
  id: string;
  name: string;
  slug: string;
  created_at: string;
}

export interface User {
  id: string;
  email: string;
  name: string;
  avatar_url?: string;
  bio?: string;
  role: "admin" | "author";
  created_at: string;
  updated_at: string;
}

// API レスポンス型
export interface PaginatedResponse<T> {
  data: T[];
  pagination: {
    page: number;
    limit: number;
    total: number;
    totalPages: number;
    hasNext: boolean;
    hasPrev: boolean;
  };
}

export type ArticleListResponse = PaginatedResponse<Article>;
