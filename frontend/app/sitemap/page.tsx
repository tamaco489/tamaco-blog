"use client";

import Link from "next/link";

export default function Sitemap() {
  const sitePages = [
    {
      category: "メインページ",
      pages: [
        { title: "ホーム", url: "/", description: "トップページ" },
        { title: "About", url: "/about", description: "サイト・運営者について" },
        { title: "お問い合わせ", url: "/contact", description: "お問い合わせフォーム" },
        { title: "ポートフォリオ", url: "/portfolio", description: "作品・プロジェクト紹介" },
        { title: "技術スタック", url: "/tech-stack", description: "使用技術・ツール" },
      ],
    },
    {
      category: "記事・コンテンツ",
      pages: [
        { title: "アーカイブ", url: "/archive", description: "記事アーカイブ（今後実装予定）" },
      ],
    },
    {
      category: "法的ページ",
      pages: [
        { title: "プライバシーポリシー", url: "/privacy", description: "個人情報の取り扱い" },
        { title: "サイトマップ", url: "/sitemap", description: "このページ" },
      ],
    },
  ];

  return (
    <div className="min-h-[calc(100vh-7rem)] bg-gray-50">
      <div className="max-w-4xl mx-auto px-4 py-8">
        <header className="text-center mb-12">
          <h1 className="text-4xl font-bold text-gray-900 mb-4">サイトマップ</h1>
          <p className="text-lg text-gray-600">当サイトの全ページ一覧</p>
        </header>

        <div className="bg-white rounded-lg shadow-sm border border-gray-200 p-8">
          <div className="space-y-10">
            {sitePages.map((section, index) => (
              <div key={index}>
                <h2 className="text-2xl font-semibold text-gray-900 mb-6 pb-2 border-b border-gray-200">
                  {section.category}
                </h2>
                <div className="grid gap-4 md:grid-cols-2">
                  {section.pages.map((page, pageIndex) => (
                    <div
                      key={pageIndex}
                      className="p-4 border border-gray-200 rounded-lg hover:border-blue-300 hover:shadow-sm transition-all duration-200"
                    >
                      <Link
                        href={page.url}
                        className="block group"
                      >
                        <h3 className="text-lg font-semibold text-blue-600 group-hover:text-blue-700 mb-2">
                          {page.title}
                        </h3>
                        <p className="text-gray-600 text-sm mb-2">
                          {page.description}
                        </p>
                        <div className="text-xs text-gray-400 font-mono">
                          {page.url}
                        </div>
                      </Link>
                    </div>
                  ))}
                </div>
              </div>
            ))}

            <div className="border-t border-gray-200 pt-8">
              <h2 className="text-2xl font-semibold text-gray-900 mb-6">
                今後実装予定のページ
              </h2>
              <div className="space-y-3 text-gray-600">
                <div className="flex items-center space-x-2">
                  <span className="w-2 h-2 bg-yellow-400 rounded-full"></span>
                  <span>記事一覧・詳細ページ（/articles）</span>
                </div>
                <div className="flex items-center space-x-2">
                  <span className="w-2 h-2 bg-yellow-400 rounded-full"></span>
                  <span>カテゴリ別記事ページ（/categories）</span>
                </div>
                <div className="flex items-center space-x-2">
                  <span className="w-2 h-2 bg-yellow-400 rounded-full"></span>
                  <span>タグ別記事ページ（/tags）</span>
                </div>
                <div className="flex items-center space-x-2">
                  <span className="w-2 h-2 bg-yellow-400 rounded-full"></span>
                  <span>検索機能（/search）</span>
                </div>
                <div className="flex items-center space-x-2">
                  <span className="w-2 h-2 bg-yellow-400 rounded-full"></span>
                  <span>管理画面（/admin）</span>
                </div>
              </div>
            </div>

            <div className="bg-blue-50 p-6 rounded-lg">
              <h3 className="text-lg font-semibold text-blue-900 mb-3">
                🛠️ 開発状況について
              </h3>
              <p className="text-blue-800 text-sm leading-relaxed">
                tamaco-blogは現在開発中のプロジェクトです。
                Next.js 14 + TypeScript + TailwindCSSを使用して構築されており、
                今後記事管理機能やCMS機能が追加される予定です。
                最新の開発状況は
                <Link href="/about" className="underline hover:text-blue-600 mx-1">
                  Aboutページ
                </Link>
                でご確認いただけます。
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}