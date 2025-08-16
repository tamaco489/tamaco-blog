"use client";


export default function TechStack() {

  const techStacks = {
    frontend: [
      {
        name: "Next.js",
        version: "14.x",
        category: "Framework",
        description: "React ベースのフルスタックフレームワーク",
        experience: "使用中",
        proficiency: 85,
        features: [
          "App Router",
          "Server Components",
          "Static Generation",
          "API Routes",
        ],
      },
      {
        name: "React",
        version: "18.x",
        category: "Library",
        description: "ユーザーインターフェース構築ライブラリ",
        experience: "使用中",
        proficiency: 80,
        features: [
          "Hooks",
          "Context API",
          "State Management",
          "Component Design",
        ],
      },
      {
        name: "TypeScript",
        version: "5.x",
        category: "Language",
        description: "型安全性を提供する JavaScript のスーパーセット",
        experience: "使用中",
        proficiency: 75,
        features: ["Type Safety", "Interface", "Generics", "Strict Mode"],
      },
      {
        name: "TailwindCSS",
        version: "3.x",
        category: "Styling",
        description: "ユーティリティファーストの CSS フレームワーク",
        experience: "使用中",
        proficiency: 90,
        features: [
          "Responsive Design",
          "Dark Mode",
          "Custom Components",
          "JIT Compiler",
        ],
      },
    ],
    backend: [
      {
        name: "Go",
        version: "1.21+",
        category: "Language",
        description: "高性能・シンプル・並行処理に優れた言語",
        experience: "学習中",
        proficiency: 40,
        features: [
          "Goroutines",
          "Channels",
          "Standard Library",
          "Cross Platform",
        ],
      },
      {
        name: "Node.js",
        version: "20.x",
        category: "Runtime",
        description: "JavaScript サーバーサイド実行環境",
        experience: "使用経験あり",
        proficiency: 65,
        features: ["Event Loop", "npm Ecosystem", "Express.js", "APIs"],
      },
      {
        name: "Python",
        version: "3.11+",
        category: "Language",
        description: "汎用プログラミング言語",
        experience: "使用経験あり",
        proficiency: 55,
        features: [
          "Data Science",
          "Web Development",
          "Automation",
          "Machine Learning",
        ],
      },
    ],
    database: [
      {
        name: "PostgreSQL",
        version: "15+",
        category: "Database",
        description: "高機能リレーショナルデータベース",
        experience: "学習中",
        proficiency: 50,
        features: [
          "ACID Transactions",
          "JSON Support",
          "Advanced Queries",
          "Performance",
        ],
      },
      {
        name: "MySQL",
        version: "8.0+",
        category: "Database",
        description: "世界で最も使用されているオープンソースデータベース",
        experience: "使用経験あり",
        proficiency: 60,
        features: ["Performance", "Replication", "Clustering", "Security"],
      },
      {
        name: "DynamoDB",
        version: "Current",
        category: "Database",
        description: "AWS のサーバーレス NoSQL データベース",
        experience: "学習中",
        proficiency: 30,
        features: ["Serverless", "Auto Scaling", "Global Tables", "Streams"],
      },
    ],
    infrastructure: [
      {
        name: "AWS",
        version: "Current",
        category: "Cloud",
        description: "Amazon Web Services クラウドプラットフォーム",
        experience: "学習中",
        proficiency: 35,
        features: ["S3", "CloudFront", "EC2", "Lambda"],
      },
      {
        name: "Docker",
        version: "24.x",
        category: "Containerization",
        description: "アプリケーションコンテナ化プラットフォーム",
        experience: "学習中",
        proficiency: 45,
        features: [
          "Containerization",
          "Multi-stage Builds",
          "Docker Compose",
          "Development Environment",
        ],
      },
      {
        name: "Terraform",
        version: "1.6+",
        category: "IaC",
        description: "Infrastructure as Code ツール",
        experience: "学習中",
        proficiency: 25,
        features: [
          "Declarative",
          "State Management",
          "Provider System",
          "Plan & Apply",
        ],
      },
    ],
    devops: [
      {
        name: "GitHub Actions",
        version: "Current",
        category: "CI/CD",
        description: "GitHub 統合 CI/CD プラットフォーム",
        experience: "使用中",
        proficiency: 70,
        features: [
          "Automated Testing",
          "Deployment",
          "Security Scanning",
          "Workflow Automation",
        ],
      },
      {
        name: "Jest",
        version: "29.x",
        category: "Testing",
        description: "JavaScript テスティングフレームワーク",
        experience: "使用中",
        proficiency: 60,
        features: [
          "Unit Testing",
          "React Testing Library",
          "Coverage Reports",
          "Mocking",
        ],
      },
    ],
    tools: [
      {
        name: "VS Code",
        version: "Current",
        category: "Editor",
        description: "Microsoft が開発したコードエディタ",
        experience: "使用中",
        proficiency: 95,
        features: [
          "Extensions",
          "Debugging",
          "Git Integration",
          "IntelliSense",
        ],
      },
      {
        name: "Git",
        version: "2.x",
        category: "VCS",
        description: "分散バージョン管理システム",
        experience: "使用中",
        proficiency: 80,
        features: [
          "Branching",
          "Merging",
          "Remote Repositories",
          "History Management",
        ],
      },
    ],
  };

  const getProficiencyColor = (proficiency: number) => {
    if (proficiency >= 80) return "bg-green-500";
    if (proficiency >= 60) return "bg-blue-500";
    if (proficiency >= 40) return "bg-yellow-500";
    return "bg-gray-400";
  };

  const getExperienceColor = (experience: string) => {
    switch (experience) {
      case "使用中":
        return "bg-green-100 text-green-800";
      case "使用経験あり":
        return "bg-blue-100 text-blue-800";
      case "学習中":
        return "bg-yellow-100 text-yellow-800";
      default:
        return "bg-gray-100 text-gray-800";
    }
  };

  const renderTechSection = (
    title: string,
    emoji: string,
    techs: typeof techStacks.frontend
  ) => (
    <div className="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <h2 className="text-2xl font-semibold text-gray-900 mb-6 flex items-center">
        <span className="mr-2">{emoji}</span>
        {title}
      </h2>

      <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
        {techs.map((tech) => (
          <div
            key={tech.name}
            className="border border-gray-200 rounded-lg p-4"
          >
            <div className="flex items-start justify-between mb-3">
              <div>
                <h3 className="text-lg font-medium text-gray-900">
                  {tech.name}
                </h3>
                <p className="text-sm text-gray-600">
                  {tech.category} • {tech.version}
                </p>
              </div>
              <span
                className={`px-2 py-1 rounded-full text-xs font-medium ${getExperienceColor(
                  tech.experience
                )}`}
              >
                {tech.experience}
              </span>
            </div>

            <p className="text-sm text-gray-700 mb-3">{tech.description}</p>

            <div className="mb-3">
              <div className="flex items-center justify-between text-sm mb-1">
                <span className="text-gray-600">習熟度</span>
                <span className="font-medium">{tech.proficiency}%</span>
              </div>
              <div className="w-full bg-gray-200 rounded-full h-2">
                <div
                  className={`h-2 rounded-full transition-all duration-300 ${getProficiencyColor(
                    tech.proficiency
                  )}`}
                  style={{ width: `${tech.proficiency}%` }}
                />
              </div>
            </div>

            <div>
              <h4 className="text-sm font-medium text-gray-900 mb-2">
                主な特徴・機能
              </h4>
              <div className="flex flex-wrap gap-1">
                {tech.features.map((feature) => (
                  <span
                    key={feature}
                    className="px-2 py-1 bg-gray-100 text-gray-700 text-xs rounded"
                  >
                    {feature}
                  </span>
                ))}
              </div>
            </div>
          </div>
        ))}
      </div>
    </div>
  );

  return (
    <div className="min-h-[calc(100vh-7rem)] bg-gray-50">
      <div className="max-w-6xl mx-auto px-4 py-8">
        <header className="text-center mb-12">
          <h1 className="text-4xl font-bold text-gray-900 mb-4">Tech Stack</h1>
          <p className="text-lg text-gray-600">使用技術・スキルセット</p>
        </header>

        <div className="space-y-8">
          {renderTechSection("Frontend", "💻", techStacks.frontend)}
          {renderTechSection("Backend", "⚙️", techStacks.backend)}
          {renderTechSection("Database", "🗄️", techStacks.database)}
          {renderTechSection("Infrastructure", "☁️", techStacks.infrastructure)}
          {renderTechSection("DevOps", "🚀", techStacks.devops)}
          {renderTechSection("Tools", "🔧", techStacks.tools)}
        </div>

        {/* 学習方針・今後の予定 */}
        <div className="mt-12 bg-white rounded-lg shadow-sm border border-gray-200 p-8">
          <h2 className="text-2xl font-semibold text-gray-900 mb-6">
            🎯 学習方針・今後の予定
          </h2>

          <div className="grid grid-cols-1 md:grid-cols-2 gap-8">
            <div>
              <h3 className="text-lg font-medium text-gray-900 mb-4">
                現在の重点分野
              </h3>
              <ul className="space-y-3 text-sm text-gray-700">
                <li className="flex items-start">
                  <span className="w-2 h-2 bg-blue-500 rounded-full mt-2 mr-3 flex-shrink-0" />
                  <div>
                    <strong>Go 言語:</strong> バックエンド API 開発の習得
                  </div>
                </li>
                <li className="flex items-start">
                  <span className="w-2 h-2 bg-blue-500 rounded-full mt-2 mr-3 flex-shrink-0" />
                  <div>
                    <strong>AWS:</strong> クラウドインフラの構築・運用
                  </div>
                </li>
                <li className="flex items-start">
                  <span className="w-2 h-2 bg-blue-500 rounded-full mt-2 mr-3 flex-shrink-0" />
                  <div>
                    <strong>Docker:</strong> コンテナ化・開発環境構築
                  </div>
                </li>
                <li className="flex items-start">
                  <span className="w-2 h-2 bg-blue-500 rounded-full mt-2 mr-3 flex-shrink-0" />
                  <div>
                    <strong>PostgreSQL:</strong> データベース設計・最適化
                  </div>
                </li>
              </ul>
            </div>

            <div>
              <h3 className="text-lg font-medium text-gray-900 mb-4">
                今後の学習計画
              </h3>
              <ul className="space-y-3 text-sm text-gray-700">
                <li className="flex items-start">
                  <span className="w-2 h-2 bg-green-500 rounded-full mt-2 mr-3 flex-shrink-0" />
                  <div>
                    <strong>短期:</strong> Go でのブログ API 完成
                  </div>
                </li>
                <li className="flex items-start">
                  <span className="w-2 h-2 bg-green-500 rounded-full mt-2 mr-3 flex-shrink-0" />
                  <div>
                    <strong>中期:</strong> AWS を使った本格運用
                  </div>
                </li>
                <li className="flex items-start">
                  <span className="w-2 h-2 bg-green-500 rounded-full mt-2 mr-3 flex-shrink-0" />
                  <div>
                    <strong>長期:</strong> マイクロサービス・分散システム
                  </div>
                </li>
                <li className="flex items-start">
                  <span className="w-2 h-2 bg-green-500 rounded-full mt-2 mr-3 flex-shrink-0" />
                  <div>
                    <strong>継続:</strong> 新しい技術の調査・検証
                  </div>
                </li>
              </ul>
            </div>
          </div>

          <div className="mt-6 p-4 bg-blue-50 rounded-lg">
            <p className="text-sm text-blue-800">
              💡 <strong>学習アプローチ:</strong>
              実際にプロジェクトを構築しながら学習することで、理論と実践を両立させています。
              このブログサイト自体も学習・実験プロジェクトとして機能しています。
            </p>
          </div>
        </div>

        {/* このサイトの技術構成 */}
        <div className="mt-8 bg-white rounded-lg shadow-sm border border-gray-200 p-8">
          <h2 className="text-2xl font-semibold text-gray-900 mb-6">
            🏗️ このサイトの技術構成
          </h2>

          <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
            <div className="text-center p-4 bg-gray-50 rounded-lg">
              <h3 className="font-medium text-gray-900 mb-2">Current (v1.0)</h3>
              <ul className="text-sm text-gray-600 space-y-1">
                <li>Next.js 14 + TypeScript</li>
                <li>TailwindCSS</li>
                <li>Mock Data</li>
                <li>GitHub Actions</li>
              </ul>
            </div>

            <div className="text-center p-4 bg-blue-50 rounded-lg">
              <h3 className="font-medium text-gray-900 mb-2">Planned (v2.0)</h3>
              <ul className="text-sm text-gray-600 space-y-1">
                <li>Go API Server</li>
                <li>PostgreSQL</li>
                <li>Docker</li>
                <li>Real CMS</li>
              </ul>
            </div>

            <div className="text-center p-4 bg-green-50 rounded-lg">
              <h3 className="font-medium text-gray-900 mb-2">Future (v3.0)</h3>
              <ul className="text-sm text-gray-600 space-y-1">
                <li>AWS Deployment</li>
                <li>CDN + S3</li>
                <li>Monitoring</li>
                <li>Performance</li>
              </ul>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
