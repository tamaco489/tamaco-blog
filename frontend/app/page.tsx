"use client";

import ArticleCard from "@/components/ArticleCard";
import Pagination from "@/components/Pagination";
import { getMockArticles, mockCategories, mockTags } from "@/lib/mock";
import { useState } from "react";

export default function Home() {
  const [currentPage, setCurrentPage] = useState(1);
  const articlesPerPage = 10;

  const articleData = getMockArticles(currentPage, articlesPerPage);

  // 全記事から各カテゴリーの記事数をカウント
  const allArticles = getMockArticles(1, 100); // 全記事を取得
  const getCategoryArticleCount = (categoryId: string) => {
    return allArticles.data.filter(
      (article) => article.category_id === categoryId
    ).length;
  };

  // 各タグの記事数をカウント
  const getTagArticleCount = (tagId: string) => {
    return allArticles.data.filter((article) =>
      article.tags?.some((tag) => tag.id === tagId)
    ).length;
  };

  const handlePageChange = (page: number) => {
    setCurrentPage(page);
  };

  return (
    <div className="bg-gray-50 flex-1">
      <div className="max-w-6xl mx-auto px-4 py-8">
        <header className="text-center mb-12">
          {/* 何か気の利いたことを書きたいです */}
        </header>

        <div className="flex gap-8">
          <div className="flex-1">
            <div className="grid grid-cols-1 md:grid-cols-2 gap-6 items-stretch">
              {articleData.data.map((article) => (
                <div key={article.id} className="w-full">
                  <ArticleCard article={article} />
                </div>
              ))}
            </div>

            <Pagination
              currentPage={currentPage}
              totalPages={articleData.pagination.totalPages}
              onPageChange={handlePageChange}
            />
          </div>

          <aside className="w-80 hidden lg:block space-y-6">
            {/* カテゴリー */}
            <div className="bg-white border border-gray-200 rounded-lg p-6">
              <h3 className="text-lg font-semibold text-gray-900 mb-4">
                カテゴリー
              </h3>
              <div className="space-y-2">
                {mockCategories.map((category) => (
                  <div
                    key={category.id}
                    className="flex items-center justify-between p-2 rounded hover:bg-gray-50 cursor-pointer"
                  >
                    <div className="flex items-center space-x-2">
                      <div
                        className="w-3 h-3 rounded-full"
                        style={{ backgroundColor: category.color || "#6B7280" }}
                      />
                      <span className="text-sm text-gray-700">
                        {category.name}
                      </span>
                    </div>
                    <span className="text-xs text-gray-500">
                      {getCategoryArticleCount(category.id)}
                    </span>
                  </div>
                ))}
              </div>
            </div>

            {/* タグ */}
            <div className="bg-white border border-gray-200 rounded-lg p-6">
              <h3 className="text-lg font-semibold text-gray-900 mb-4">タグ</h3>
              <div className="flex flex-wrap gap-2">
                {mockTags.map((tag) => (
                  <div
                    key={tag.id}
                    className="inline-flex items-center space-x-1 px-3 py-1 bg-gray-100 hover:bg-gray-200 rounded-full cursor-pointer transition-colors"
                  >
                    <span className="text-sm text-gray-700">#{tag.name}</span>
                    <span className="text-xs text-gray-500">
                      ({getTagArticleCount(tag.id)})
                    </span>
                  </div>
                ))}
              </div>
            </div>
          </aside>
        </div>
      </div>
    </div>
  );
}
