"use client";

import { useEffect } from "react";

export default function Portfolio() {
  useEffect(() => {
    console.log("💼 Portfolio ページを表示中... tamaco-blog/portfolio");
  }, []);

  const projects = [
    {
      id: "tamaco-blog",
      title: "tamaco-blog",
      description: "Next.js + TypeScript + TailwindCSS で構築したテックブログ",
      technologies: ["Next.js 14", "TypeScript", "TailwindCSS", "Jest"],
      status: "開発中",
      github: "https://github.com/tamaco489/tamaco-blog",
      demo: null,
      features: [
        "App Router を使用したモダンな構成",
        "レスポンシブデザイン",
        "記事のカテゴリー・タグ管理",
        "ページネーション機能",
        "CI/CD パイプライン"
      ]
    },
    {
      id: "future-project-1",
      title: "Go API Server",
      description: "ブログのバックエンド API（予定）",
      technologies: ["Go", "PostgreSQL", "Docker"],
      status: "企画中",
      github: null,
      demo: null,
      features: [
        "RESTful API 設計",
        "JWT 認証",
        "データベース設計",
        "Docker コンテナ化"
      ]
    },
    {
      id: "future-project-2",
      title: "AWS Infrastructure",
      description: "クラウドインフラ構築（予定）",
      technologies: ["AWS", "Terraform", "CloudFront", "S3"],
      status: "企画中",
      github: null,
      demo: null,
      features: [
        "Infrastructure as Code",
        "CDN 配信",
        "CI/CD デプロイ",
        "セキュリティ設定"
      ]
    }
  ];

  return (
    <div className="min-h-[calc(100vh-7rem)] bg-gray-50">
      <div className="max-w-6xl mx-auto px-4 py-8">
        <header className="text-center mb-12">
          <h1 className="text-4xl font-bold text-gray-900 mb-4">Portfolio</h1>
          <p className="text-lg text-gray-600">
            これまでの制作物・取り組み
            <br />
            <span className="text-sm text-gray-500 italic">
              console.log("プロジェクト一覧を表示中💼");
            </span>
          </p>
        </header>

        <div className="grid grid-cols-1 lg:grid-cols-2 xl:grid-cols-3 gap-8">
          {projects.map((project) => (
            <div
              key={project.id}
              className="bg-white rounded-lg shadow-sm border border-gray-200 p-6 hover:shadow-md transition-shadow"
            >
              <div className="flex items-start justify-between mb-4">
                <h2 className="text-xl font-semibold text-gray-900">
                  {project.title}
                </h2>
                <span
                  className={`px-2 py-1 rounded-full text-xs font-medium ${
                    project.status === "開発中"
                      ? "bg-blue-100 text-blue-800"
                      : "bg-yellow-100 text-yellow-800"
                  }`}
                >
                  {project.status}
                </span>
              </div>

              <p className="text-gray-600 mb-4 text-sm leading-relaxed">
                {project.description}
              </p>

              <div className="mb-4">
                <h3 className="text-sm font-medium text-gray-900 mb-2">技術スタック</h3>
                <div className="flex flex-wrap gap-1">
                  {project.technologies.map((tech) => (
                    <span
                      key={tech}
                      className="px-2 py-1 bg-gray-100 text-gray-700 text-xs rounded"
                    >
                      {tech}
                    </span>
                  ))}
                </div>
              </div>

              <div className="mb-4">
                <h3 className="text-sm font-medium text-gray-900 mb-2">主な機能</h3>
                <ul className="text-xs text-gray-600 space-y-1">
                  {project.features.map((feature, index) => (
                    <li key={index} className="flex items-start">
                      <span className="w-1 h-1 bg-gray-400 rounded-full mt-2 mr-2 flex-shrink-0" />
                      {feature}
                    </li>
                  ))}
                </ul>
              </div>

              <div className="flex space-x-3 mt-6">
                {project.github && (
                  <a
                    href={project.github}
                    target="_blank"
                    rel="noopener noreferrer"
                    className="flex-1 bg-gray-900 hover:bg-gray-800 text-white text-sm font-medium py-2 px-4 rounded-lg transition-colors text-center"
                  >
                    GitHub
                  </a>
                )}
                {project.demo && (
                  <a
                    href={project.demo}
                    target="_blank"
                    rel="noopener noreferrer"
                    className="flex-1 bg-blue-600 hover:bg-blue-700 text-white text-sm font-medium py-2 px-4 rounded-lg transition-colors text-center"
                  >
                    Demo
                  </a>
                )}
                {!project.github && !project.demo && (
                  <div className="flex-1 bg-gray-100 text-gray-500 text-sm font-medium py-2 px-4 rounded-lg text-center">
                    準備中
                  </div>
                )}
              </div>
            </div>
          ))}
        </div>

        <div className="mt-12 text-center">
          <div className="bg-white rounded-lg shadow-sm border border-gray-200 p-8">
            <h2 className="text-2xl font-semibold text-gray-900 mb-4">
              スキル・経験
            </h2>
            
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 text-left">
              <div>
                <h3 className="font-medium text-gray-900 mb-3">フロントエンド</h3>
                <ul className="text-sm text-gray-600 space-y-1">
                  <li>• React / Next.js</li>
                  <li>• TypeScript / JavaScript</li>
                  <li>• TailwindCSS / CSS</li>
                  <li>• Jest / Testing Library</li>
                </ul>
              </div>

              <div>
                <h3 className="font-medium text-gray-900 mb-3">バックエンド</h3>
                <ul className="text-sm text-gray-600 space-y-1">
                  <li>• Go（学習中）</li>
                  <li>• Node.js</li>
                  <li>• PostgreSQL</li>
                  <li>• RESTful API</li>
                </ul>
              </div>

              <div>
                <h3 className="font-medium text-gray-900 mb-3">インフラ・ツール</h3>
                <ul className="text-sm text-gray-600 space-y-1">
                  <li>• AWS（学習中）</li>
                  <li>• Docker</li>
                  <li>• GitHub Actions</li>
                  <li>• Git / GitHub</li>
                </ul>
              </div>
            </div>

            <div className="mt-6 p-4 bg-gray-50 rounded-lg">
              <p className="text-sm text-gray-600 italic">
                💡 現在学習・実践を通じてスキルアップに取り組んでいます。
                新しい技術への挑戦と継続的な成長を大切にしています。
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}