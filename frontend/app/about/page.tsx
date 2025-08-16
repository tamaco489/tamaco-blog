"use client";


export default function About() {

  return (
    <div className="min-h-[calc(100vh-7rem)] bg-gray-50">
      <div className="max-w-4xl mx-auto px-4 py-8">
        <header className="text-center mb-12">
          <h1 className="text-4xl font-bold text-gray-900 mb-4">About</h1>
          <p className="text-lg text-gray-600">このブログについて</p>
        </header>

        <div className="bg-white rounded-lg shadow-sm border border-gray-200 p-8">
          <div className="prose prose-lg max-w-none">
            <h2 className="text-2xl font-semibold text-gray-900 mb-6">
              tamaco-blog について
            </h2>

            <div className="space-y-6 text-gray-700">
              <p>
                画面の向こうから、サーバーの奥まで 🚀
                <br />
                技術的な発見や学びを共有するテックブログです。
              </p>

              <div>
                <h3 className="text-xl font-semibold text-gray-900 mb-3">
                  このブログで扱うトピック
                </h3>
                <ul className="list-disc list-inside space-y-2">
                  <li>フロントエンド開発（React, Next.js, TypeScript）</li>
                  <li>バックエンド開発（Go, Node.js）</li>
                  <li>インフラ・クラウド（AWS, Docker）</li>
                  <li>開発ツール・環境構築</li>
                  <li>プログラミング Tips & Tricks</li>
                </ul>
              </div>

              <div>
                <h3 className="text-xl font-semibold text-gray-900 mb-3">
                  技術スタック
                </h3>
                <p>このブログは以下の技術で構築されています：</p>
                <ul className="list-disc list-inside space-y-2 mt-2">
                  <li>
                    <strong>フロントエンド:</strong> Next.js 14, TypeScript,
                    TailwindCSS
                  </li>
                  <li>
                    <strong>バックエンド:</strong> Go (予定)
                  </li>
                  <li>
                    <strong>データベース:</strong> Supabase PostgreSQL (予定)
                  </li>
                  <li>
                    <strong>ホスティング:</strong> AWS S3 + CloudFront (予定)
                  </li>
                  <li>
                    <strong>CI/CD:</strong> GitHub Actions
                  </li>
                </ul>
              </div>

              <div className="bg-gray-50 p-4 rounded-lg">
                <p className="text-sm text-gray-600 italic">
                  💡
                  このサイトはモダンな技術スタックを使用した学習・実験プロジェクトとしても機能しています。
                </p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
