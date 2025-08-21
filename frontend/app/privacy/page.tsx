"use client";

export default function Privacy() {
  return (
    <div className="min-h-[calc(100vh-7rem)] bg-gray-50">
      <div className="max-w-4xl mx-auto px-4 py-8">
        <header className="text-center mb-12">
          <h1 className="text-4xl font-bold text-gray-900 mb-4">プライバシーポリシー</h1>
          <p className="text-lg text-gray-600">アクセス解析とCookieの使用について</p>
        </header>

        <div className="bg-white rounded-lg shadow-sm border border-gray-200 p-8">
          <div className="prose prose-lg max-w-none">
            <div className="space-y-8 text-gray-700">
              <div>
                <h2 className="text-2xl font-semibold text-gray-900 mb-4">
                  アクセス解析ツールについて
                </h2>
                <p>
                  当サイトでは、サイトの利用状況を把握し、より良いサービスを提供するために以下のツールを使用しています。
                  これらのツールは、匿名化されたデータを収集しており、個人を特定することはありません。
                </p>
                <ul className="list-disc list-inside space-y-2 mt-4">
                  <li><strong>Google Analytics</strong>: アクセス状況の分析</li>
                  <li><strong>New Relic</strong>: アプリケーションの監視・パフォーマンス測定</li>
                </ul>
              </div>

              <div>
                <h2 className="text-2xl font-semibold text-gray-900 mb-4">
                  Cookieの使用について
                </h2>
                <p>
                  当サイトでは、アクセス解析やパフォーマンス向上のためにCookieを使用しています。
                  Cookieは以下の目的で使用されます：
                </p>
                <ul className="list-disc list-inside space-y-2 mt-4">
                  <li>サイトの利用状況の分析</li>
                  <li>ユーザーエクスペリエンスの向上</li>
                  <li>サイトのパフォーマンス測定</li>
                </ul>
                <p className="mt-4">
                  Cookieの使用を無効にしたい場合は、ブラウザの設定から変更することができます。
                  ただし、一部の機能が制限される場合があります。
                </p>
              </div>

              <div>
                <h2 className="text-2xl font-semibold text-gray-900 mb-4">
                  お問い合わせについて
                </h2>
                <p>
                  お問い合わせフォームをご利用いただく際は、お名前とメールアドレスをお預かりします。
                  これらの情報は、お問い合わせへの回答にのみ使用し、それ以外の目的では使用いたしません。
                </p>
              </div>

              <div>
                <h2 className="text-2xl font-semibold text-gray-900 mb-4">
                  ポリシーの変更
                </h2>
                <p>
                  このプライバシーポリシーは、必要に応じて変更される場合があります。
                  変更後のポリシーは、当サイトに掲載された時点で効力を生じます。
                </p>
              </div>

              <div className="bg-blue-50 p-6 rounded-lg">
                <h3 className="text-lg font-semibold text-blue-900 mb-3">
                  📋 要約
                </h3>
                <p className="text-blue-800 text-sm leading-relaxed">
                  当サイトは技術ブログとして運営されており、閲覧に際して個人情報の登録は不要です。
                  アクセス解析のためにCookieを使用していますが、これらは匿名化されており、
                  個人を特定する情報は収集していません。
                </p>
              </div>

              <div className="bg-gray-50 p-4 rounded-lg">
                <p className="text-sm text-gray-600">
                  最終更新日: 2024年8月16日
                  <br />
                  運営者: tamaco489
                  <br />
                  お問い合わせ: このポリシーに関するご質問は、お問い合わせページよりご連絡ください。
                </p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
