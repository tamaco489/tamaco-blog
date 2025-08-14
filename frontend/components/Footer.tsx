import Link from "next/link";

/**
 * Footer コンポーネント
 *
 * 目的:
 * - サイト下部に GitHub 風のフッターを設置する。
 *   - 左側に © 年 とサイト名
 *   - 右側に水平ナビゲーションリンク（必要最低限）
 *
 * 実装内容:
 * - 上部にボーダーを表示し、横幅は全幅。内部は最大幅で中央寄せ。
 * - 小画面では縦積み、広い画面では左右に分割されるレイアウト。
 *
 * 影響範囲:
 * - `app/layout.tsx` に取り込み、全ページに表示される。
 */
export function Footer() {
  const year = new Date().getFullYear();
  return (
    <footer className="w-full border-t bg-white mt-32">
      <div className="mx-auto max-w-screen-lg px-4 py-8">
        <div className="flex flex-col md:flex-row items-center md:items-start justify-between gap-4">
          <p className="text-sm text-gray-600">© {year} tamaco-blog</p>

          <nav aria-label="フッターナビゲーション" className="text-sm text-gray-600">
            <ul className="flex flex-wrap items-center gap-x-4 gap-y-2">
              <li>
                <Link href="/about" className="hover:underline">
                  About
                </Link>
              </li>
              <li>
                <Link href="/docs" className="hover:underline">
                  Docs
                </Link>
              </li>
              <li>
                <a
                  href="https://github.com/tamaco489/tamaco-blog"
                  target="_blank"
                  rel="noopener noreferrer"
                  className="hover:underline"
                >
                  GitHub
                </a>
              </li>
              <li>
                <a
                  href="https://www.githubstatus.com/"
                  target="_blank"
                  rel="noopener noreferrer"
                  className="hover:underline"
                >
                  Status
                </a>
              </li>
            </ul>
          </nav>
        </div>
      </div>
    </footer>
  );
}

export default Footer;


