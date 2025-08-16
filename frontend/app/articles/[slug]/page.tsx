"use client";

import { getMockArticleBySlug } from "@/lib/mock";
import Link from "next/link";
import { notFound } from "next/navigation";

interface ArticleDetailPageProps {
  params: {
    slug: string;
  };
}

export default function ArticleDetailPage({ params }: ArticleDetailPageProps) {
  const article = getMockArticleBySlug(params.slug);

  if (!article) {
    notFound();
  }

  return (
    <div className="min-h-[calc(100vh-7rem)] bg-gray-50">
      <div className="max-w-4xl mx-auto px-4 py-8">
        {/* パンくずリスト */}
        <nav className="flex items-center space-x-2 text-sm text-gray-600 mb-8">
          <Link href="/" className="hover:text-blue-600 transition-colors">
            Home
          </Link>
          <span>/</span>
          <span className="text-gray-900">記事詳細</span>
        </nav>

        {/* 記事ヘッダー */}
        <header className="mb-8">
          <h1 className="text-4xl font-bold text-gray-900 mb-4 leading-tight">
            {article.title}
          </h1>

          <div className="flex items-center justify-between flex-wrap gap-4 mb-6">
            <div className="flex items-center space-x-4">
              {article.category && (
                <span
                  className="px-3 py-1 rounded-full text-white text-sm font-medium"
                  style={{
                    backgroundColor: article.category.color || "#6B7280",
                  }}
                >
                  {article.category.name}
                </span>
              )}
              <span className="text-gray-600 text-sm">
                {new Date(
                  article.published_at || article.created_at
                ).toLocaleDateString("ja-JP", {
                  year: "numeric",
                  month: "long",
                  day: "numeric",
                })}
              </span>
            </div>

            <div className="flex items-center space-x-4 text-sm text-gray-600">
              <span>{article.view_count.toLocaleString()} views</span>
              {article.author && (
                <div className="flex items-center space-x-2">
                  <span>by</span>
                  <span className="font-medium text-gray-900">
                    {article.author.name}
                  </span>
                </div>
              )}
            </div>
          </div>

          {/* タグ */}
          {article.tags && article.tags.length > 0 && (
            <div className="flex flex-wrap gap-2 mb-6">
              {article.tags.map((tag) => (
                <span
                  key={tag.id}
                  className="px-3 py-1 bg-gray-100 text-gray-700 rounded-full text-sm hover:bg-gray-200 transition-colors cursor-pointer"
                >
                  #{tag.name}
                </span>
              ))}
            </div>
          )}

          {/* 記事の概要 */}
          {article.excerpt && (
            <div className="bg-blue-50 border-l-4 border-blue-400 p-4 mb-8">
              <p className="text-gray-700 leading-relaxed">{article.excerpt}</p>
            </div>
          )}
        </header>

        {/* 記事本文 */}
        <article className="bg-white rounded-lg shadow-sm border border-gray-200 p-8 mb-8">
          <div className="prose prose-lg max-w-none">
            <div className="text-gray-700 leading-relaxed whitespace-pre-wrap">
              {article.content}
            </div>
          </div>
        </article>

        {/* 記事フッター */}
        <footer className="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <div className="flex items-center justify-between">
            <div>
              <p className="text-sm text-gray-600 mb-2">最終更新</p>
              <p className="text-gray-900">
                {new Date(article.updated_at).toLocaleDateString("ja-JP", {
                  year: "numeric",
                  month: "long",
                  day: "numeric",
                  hour: "2-digit",
                  minute: "2-digit",
                })}
              </p>
            </div>

            <Link
              href="/"
              className="inline-flex items-center px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors"
            >
              記事一覧に戻る
            </Link>
          </div>
        </footer>
      </div>
    </div>
  );
}
