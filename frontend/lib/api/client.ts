import { Article, ArticleListResponse, Category, Tag } from "@/types/blog";

const API_BASE_URL = process.env.NEXT_PUBLIC_API_BASE_URL || "";

class ApiClient {
  private baseUrl: string;

  constructor(baseUrl: string = API_BASE_URL) {
    this.baseUrl = baseUrl;
  }

  private async request<T>(
    endpoint: string,
    options?: RequestInit
  ): Promise<T> {
    const url = `${this.baseUrl}${endpoint}`;

    const defaultOptions: RequestInit = {
      headers: {
        "Content-Type": "application/json",
        Accept: "application/json",
        "User-Agent": "tamaco-blog/1.0",
        "X-Requested-With": "XMLHttpRequest",
        "Cache-Control": "no-cache",
      },
    };

    const config = { ...defaultOptions, ...options };

    try {
      const response = await fetch(url, config);

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }

      return await response.json();
    } catch (error) {
      console.error(`API request failed for ${endpoint}:`, error);
      throw error;
    }
  }

  // 記事関連API
  async getArticles(
    page: number = 1,
    limit: number = 10
  ): Promise<ArticleListResponse> {
    const params = new URLSearchParams({
      page: page.toString(),
      limit: limit.toString(),
    });

    return this.request<ArticleListResponse>(`/api/articles?${params}`);
  }

  async getArticleBySlug(slug: string): Promise<Article> {
    return this.request<Article>(`/api/articles/${slug}`);
  }

  // カテゴリ関連API
  async getCategories(): Promise<Category[]> {
    return this.request<Category[]>("/api/categories");
  }

  // タグ関連API
  async getTags(): Promise<Tag[]> {
    return this.request<Tag[]>("/api/tags");
  }
}

// シングルトンインスタンスをエクスポート
export const apiClient = new ApiClient();

// 個別の関数もエクスポート（使いやすさのため）
export const fetchArticles = (page?: number, limit?: number) =>
  apiClient.getArticles(page, limit);

export const fetchArticleBySlug = (slug: string) =>
  apiClient.getArticleBySlug(slug);

export const fetchCategories = () => apiClient.getCategories();

export const fetchTags = () => apiClient.getTags();
