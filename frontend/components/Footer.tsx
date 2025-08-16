import Image from "next/image";
import Link from "next/link";

/**
 * Footer コンポーネント
 *
 * 目的:
 * - サイト下部に充実したフッターを設置する。
 *   - サイト情報（ロゴ、説明、コピーライト）
 *   - 4列のナビゲーション（サイト、プロフィール、その他）
 *   - 技術スタック情報
 *
 * 実装内容:
 * - 4列グリッドレイアウト（レスポンシブ）
 * - Headerメニューとの統一感
 * - 外部リンクに外部リンクアイコン
 *
 * 影響範囲:
 * - `app/layout.tsx` に取り込み、全ページに表示される。
 */
export function Footer() {
  const year = new Date().getFullYear();
  return (
    <footer className="w-full border-t border-gray-700 bg-gray-900 mt-16 relative z-10">
      <div className="mx-auto max-w-6xl px-4 py-12">
        <div className="grid grid-cols-1 md:grid-cols-4 gap-8">
          {/* サイト情報 */}
          <div className="md:col-span-1">
            <div className="flex items-center gap-3 mb-4">
              <Image
                src="/blog/icon/favicon.png"
                alt="tamaco-blog"
                width={32}
                height={32}
                className="rounded-lg"
              />
              <span className="text-xl font-semibold text-white">
                tamaco-blog
              </span>
            </div>
            <p className="text-sm text-gray-400 mb-4">
              フロントエンドエンジニアの技術ブログ。Next.js、React、TypeScriptなどの技術情報を発信しています。
            </p>
            {/* <p className="text-xs text-gray-500">© {year} tamaco-blog</p> */}
          </div>

          {/* サイトナビゲーション */}
          <div className="md:col-span-1">
            <h3 className="text-sm font-semibold text-white mb-4">サイト</h3>
            <ul className="space-y-2">
              <li>
                <Link
                  href="/"
                  className="text-sm text-gray-400 hover:text-white transition-colors"
                >
                  ホーム
                </Link>
              </li>
              <li>
                <Link
                  href="/archive"
                  className="text-sm text-gray-400 hover:text-white transition-colors"
                >
                  アーカイブ
                </Link>
              </li>
              <li>
                <Link
                  href="/about"
                  className="text-sm text-gray-400 hover:text-white transition-colors"
                >
                  About
                </Link>
              </li>
              <li>
                <Link
                  href="/contact"
                  className="text-sm text-gray-400 hover:text-white transition-colors"
                >
                  Contact
                </Link>
              </li>
            </ul>
          </div>

          {/* プロフィール・ポートフォリオ */}
          <div className="md:col-span-1">
            <h3 className="text-sm font-semibold text-white mb-4">
              プロフィール
            </h3>
            <ul className="space-y-2">
              <li>
                <Link
                  href="/portfolio"
                  className="text-sm text-gray-400 hover:text-white transition-colors"
                >
                  Portfolio
                </Link>
              </li>
              <li>
                <Link
                  href="/tech-stack"
                  className="text-sm text-gray-400 hover:text-white transition-colors"
                >
                  Tech Stack
                </Link>
              </li>
              <li>
                <a
                  href="https://github.com/tamaco489"
                  target="_blank"
                  rel="noopener noreferrer"
                  className="text-sm text-gray-400 hover:text-white transition-colors inline-flex items-center gap-1"
                >
                  GitHub
                  <svg
                    className="w-3 h-3"
                    fill="currentColor"
                    viewBox="0 0 20 20"
                  >
                    <path
                      fillRule="evenodd"
                      d="M4.25 5.5a.75.75 0 00-.75.75v8.5c0 .414.336.75.75.75h8.5a.75.75 0 00.75-.75v-4a.75.75 0 011.5 0v4A2.25 2.25 0 0112.75 17h-8.5A2.25 2.25 0 012 14.75v-8.5A2.25 2.25 0 014.25 4h5a.75.75 0 010 1.5h-5z"
                      clipRule="evenodd"
                    />
                    <path
                      fillRule="evenodd"
                      d="M6.194 12.753a.75.75 0 001.06.053L16.5 4.44v2.81a.75.75 0 001.5 0v-4.5a.75.75 0 00-.75-.75h-4.5a.75.75 0 000 1.5h2.553l-9.056 8.194a.75.75 0 00-.053 1.06z"
                      clipRule="evenodd"
                    />
                  </svg>
                </a>
              </li>
            </ul>
          </div>

          {/* その他・法的情報 */}
          <div className="md:col-span-1">
            <h3 className="text-sm font-semibold text-white mb-4">その他</h3>
            <ul className="space-y-2">
              <li>
                <Link
                  href="/privacy"
                  className="text-sm text-gray-400 hover:text-white transition-colors"
                >
                  プライバシーポリシー
                </Link>
              </li>
              <li>
                <Link
                  href="/sitemap"
                  className="text-sm text-gray-400 hover:text-white transition-colors"
                >
                  サイトマップ
                </Link>
              </li>
            </ul>
          </div>
        </div>

        {/* ボトムライン */}
        <div className="border-t border-gray-700 mt-8 pt-6">
          <div className="flex flex-col sm:flex-row items-center justify-between gap-4">
            <p className="text-xs text-gray-500">© {year} tamaco-blog</p>
            <p className="text-xs text-gray-500">Built with Next.js & AWS</p>
          </div>
        </div>
      </div>
    </footer>
  );
}

export default Footer;
