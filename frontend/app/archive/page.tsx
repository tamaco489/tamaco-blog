"use client";

import { getMockArticles, mockCategories } from "@/lib/mock";
import Link from "next/link";
import { useEffect, useState } from "react";

export default function Archive() {
  const [selectedCategory, setSelectedCategory] = useState<string>("all");
  const [selectedYear, setSelectedYear] = useState<string>("all");

  useEffect(() => {
    console.log("📚 Archive ページを表示中... tamaco-blog/archive");
  }, []);

  const allArticles = getMockArticles(1, 100).data;

  // 年度別グルーピング用のデータ準備
  const years = [...new Set(allArticles.map(article => 
    new Date(article.created_at).getFullYear().toString()
  ))].sort((a, b) => parseInt(b) - parseInt(a));

  // フィルタリングされた記事
  const filteredArticles = allArticles.filter(article => {
    const categoryMatch = selectedCategory === "all" || article.category_id === selectedCategory;
    const yearMatch = selectedYear === "all" || 
      new Date(article.created_at).getFullYear().toString() === selectedYear;
    return categoryMatch && yearMatch;
  });

  // 年月別にグルーピング
  const articlesByYearMonth = filteredArticles.reduce((acc, article) => {
    const date = new Date(article.created_at);
    const yearMonth = `${date.getFullYear()}-${String(date.getMonth() + 1).padStart(2, '0')}`;
    if (!acc[yearMonth]) {
      acc[yearMonth] = [];
    }
    acc[yearMonth].push(article);
    return acc;
  }, {} as Record<string, typeof filteredArticles>);

  const sortedYearMonths = Object.keys(articlesByYearMonth).sort((a, b) => b.localeCompare(a));

  const formatYearMonth = (yearMonth: string) => {
    const [year, month] = yearMonth.split('-');
    return `${year}年${parseInt(month)}月`;
  };

  const getCategoryName = (categoryId: string) => {
    return mockCategories.find(cat => cat.id === categoryId)?.name || "Unknown";
  };

  const getCategoryColor = (categoryId: string) => {
    return mockCategories.find(cat => cat.id === categoryId)?.color || "#6B7280";
  };

  return (
    <div className="min-h-[calc(100vh-7rem)] bg-gray-50">
      <div className="max-w-6xl mx-auto px-4 py-8">
        <header className="text-center mb-12">
          <h1 className="text-4xl font-bold text-gray-900 mb-4">Archive</h1>
          <p className="text-lg text-gray-600">
            記事一覧・アーカイブ
            <br />
            <span className="text-sm text-gray-500 italic">
              console.log("記事アーカイブを表示中📚");
            </span>
          </p>
        </header>

        {/* フィルター */}
        <div className="bg-white rounded-lg shadow-sm border border-gray-200 p-6 mb-8">
          <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div>
              <label htmlFor="category-filter" className="block text-sm font-medium text-gray-700 mb-2">
                カテゴリー
              </label>
              <select
                id="category-filter"
                value={selectedCategory}
                onChange={(e) => setSelectedCategory(e.target.value)}
                className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-gray-500"
              >
                <option value="all">すべて</option>
                {mockCategories.map((category) => (
                  <option key={category.id} value={category.id}>
                    {category.name}
                  </option>
                ))}
              </select>
            </div>

            <div>
              <label htmlFor="year-filter" className="block text-sm font-medium text-gray-700 mb-2">
                年度
              </label>
              <select
                id="year-filter"
                value={selectedYear}
                onChange={(e) => setSelectedYear(e.target.value)}
                className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-gray-500"
              >
                <option value="all">すべて</option>
                {years.map((year) => (
                  <option key={year} value={year}>
                    {year}年
                  </option>
                ))}
              </select>
            </div>
          </div>
          
          <div className="mt-4 text-sm text-gray-600">
            {filteredArticles.length} 件の記事が見つかりました
          </div>
        </div>

        {/* 記事一覧 */}
        <div className="space-y-8">
          {sortedYearMonths.length === 0 ? (
            <div className="text-center py-12">
              <p className="text-gray-500">該当する記事が見つかりませんでした。</p>
            </div>
          ) : (
            sortedYearMonths.map((yearMonth) => (
              <div key={yearMonth} className="bg-white rounded-lg shadow-sm border border-gray-200">
                <div className="bg-gray-50 px-6 py-3 border-b border-gray-200">
                  <h2 className="text-lg font-semibold text-gray-900">
                    {formatYearMonth(yearMonth)}
                  </h2>
                  <p className="text-sm text-gray-600">
                    {articlesByYearMonth[yearMonth].length} 件の記事
                  </p>
                </div>
                
                <div className="p-6">
                  <div className="space-y-4">
                    {articlesByYearMonth[yearMonth].map((article) => (
                      <div
                        key={article.id}
                        className="flex items-start space-x-4 p-4 hover:bg-gray-50 rounded-lg transition-colors"
                      >
                        <div className="flex-shrink-0 text-sm text-gray-500 w-20">
                          {new Date(article.created_at).toLocaleDateString('ja-JP', {
                            month: '2-digit',
                            day: '2-digit'
                          })}
                        </div>
                        
                        <div className="flex-1 min-w-0">
                          <div className="flex items-start justify-between">
                            <div className="flex-1">
                              <Link
                                href={`/articles/${article.slug}`}
                                className="text-lg font-medium text-gray-900 hover:text-blue-600 transition-colors block"
                              >
                                {article.title}
                              </Link>
                              <p className="text-sm text-gray-600 mt-1 line-clamp-2">
                                {article.summary || article.content.substring(0, 100) + "..."}
                              </p>
                            </div>
                          </div>
                          
                          <div className="flex items-center space-x-4 mt-3">
                            {article.category && (
                              <div className="flex items-center space-x-1">
                                <div
                                  className="w-3 h-3 rounded-full"
                                  style={{ backgroundColor: getCategoryColor(article.category_id) }}
                                />
                                <span className="text-sm text-gray-600">
                                  {getCategoryName(article.category_id)}
                                </span>
                              </div>
                            )}
                            
                            {article.tags && article.tags.length > 0 && (
                              <div className="flex flex-wrap gap-1">
                                {article.tags.slice(0, 3).map((tag) => (
                                  <span
                                    key={tag.id}
                                    className="px-2 py-1 bg-gray-100 text-gray-600 text-xs rounded"
                                  >
                                    #{tag.name}
                                  </span>
                                ))}
                                {article.tags.length > 3 && (
                                  <span className="text-xs text-gray-500">
                                    +{article.tags.length - 3}
                                  </span>
                                )}
                              </div>
                            )}
                            
                            <div className="text-sm text-gray-500">
                              {article.view_count} views
                            </div>
                          </div>
                        </div>
                      </div>
                    ))}
                  </div>
                </div>
              </div>
            ))
          )}
        </div>

        {/* 統計情報 */}
        <div className="mt-12 bg-white rounded-lg shadow-sm border border-gray-200 p-6">
          <h2 className="text-xl font-semibold text-gray-900 mb-4">統計情報</h2>
          
          <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
            <div className="text-center">
              <div className="text-3xl font-bold text-gray-900">{allArticles.length}</div>
              <div className="text-sm text-gray-600">総記事数</div>
            </div>
            
            <div className="text-center">
              <div className="text-3xl font-bold text-gray-900">{mockCategories.length}</div>
              <div className="text-sm text-gray-600">カテゴリー数</div>
            </div>
            
            <div className="text-center">
              <div className="text-3xl font-bold text-gray-900">{years.length}</div>
              <div className="text-sm text-gray-600">活動年数</div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}