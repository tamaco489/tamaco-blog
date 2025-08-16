"use client";

import ArticleCard from "@/components/ArticleCard";
import Pagination from "@/components/Pagination";
import { getMockArticles } from "@/lib/mock-data";
import { useState } from "react";

export default function Home() {
  const [currentPage, setCurrentPage] = useState(1);
  const articlesPerPage = 10;

  const articleData = getMockArticles(currentPage, articlesPerPage);

  const handlePageChange = (page: number) => {
    setCurrentPage(page);
  };

  return (
    <div className="bg-gray-50 flex-1">
      <div className="max-w-6xl mx-auto px-4 py-8">
        <header className="text-center mb-12">
          <h1 className="text-4xl font-bold text-gray-900 mb-4">tamaco Blog</h1>
          <p className="text-lg text-gray-600">ブログ</p>
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

          <aside className="w-80 hidden lg:block">
            <div className="bg-white border border-gray-200 rounded-lg p-6">
              <h3 className="text-lg font-semibold text-gray-900 mb-4">
                カテゴリー・タグ
              </h3>
              <p className="text-gray-600 text-sm">
                将来的にカテゴリーやタグの一覧を表示予定
              </p>
            </div>
          </aside>
        </div>
      </div>
    </div>
  );
}
