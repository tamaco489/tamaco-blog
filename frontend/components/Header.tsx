"use client";
import Image from "next/image";
import Link from "next/link";
import { useEffect, useRef, useState } from "react";

/**
 * Header コンポーネント
 *
 * 目的:
 * - 画面上部にヘッダーを設置し、以下を表示する。
 *   - 左: MENU（ホバー/フォーカスでプロフィールやGitHubへのリンクを縦並び表示）
 *   - 中央: サイトタイトル "tamaco-blog"
 *   - 右: 検索窓（将来的に全文検索を実装予定のプレースホルダ）
 *
 * 実装内容:
 * - CSS Grid（3カラム）で左・中央・右の領域をレイアウト。
 * - MENU は group-hover / focus-within でドロップダウンを表示。
 * - 検索窓は現時点では機能しないプレースホルダの input。
 *
 * 影響範囲:
 * - `app/layout.tsx` に組み込み、全ページのヘッダーとして表示します。
 */
export function Header() {
  const [menuOpen, setMenuOpen] = useState(false);
  const menuRef = useRef<HTMLDivElement | null>(null);

  // Esc キーでメニューを閉じる
  useEffect(() => {
    function handleKeydown(event: KeyboardEvent) {
      if (event.key === "Escape") setMenuOpen(false);
    }
    document.addEventListener("keydown", handleKeydown);
    return () => {
      document.removeEventListener("keydown", handleKeydown);
    };
  }, []);

  // メニュー展開中は背景スクロールをロック
  useEffect(() => {
    const { body } = document;
    const previous = body.style.overflow;
    if (menuOpen) body.style.overflow = "hidden";
    else body.style.overflow = previous;
    return () => {
      body.style.overflow = previous;
    };
  }, [menuOpen]);

  return (
    <header className="w-full border-b border-gray-700 bg-gray-900">
      <div className="w-full h-28 grid grid-cols-3 items-center px-4">
        {/* 左: ハンバーガーメニュー */}
        <div ref={menuRef} className="relative justify-self-start">
          <button
            className="inline-flex flex-col items-center justify-center gap-1.5 h-10 w-10 rounded-md hover:bg-gray-700 focus:outline-none focus-visible:ring-2 focus-visible:ring-gray-500"
            aria-haspopup="true"
            aria-expanded={menuOpen}
            aria-label="メニュー"
            onClick={() => setMenuOpen((v) => !v)}
          >
            <span className="sr-only">メニュー</span>
            <span className="block h-0.5 w-6 bg-gray-300" />
            <span className="block h-0.5 w-6 bg-gray-300" />
            <span className="block h-0.5 w-6 bg-gray-300" />
          </button>
        </div>

        {/* スライドメニュー（左から、全高） */}
        <div
          className={`fixed inset-0 z-40 transition-opacity duration-300 bg-black/40 ${
            menuOpen
              ? "opacity-100 pointer-events-auto"
              : "opacity-0 pointer-events-none"
          }`}
          onClick={() => setMenuOpen(false)}
          aria-hidden={!menuOpen}
        />

        <nav
          role="dialog"
          aria-modal={menuOpen}
          aria-label="メニュー"
          className={`fixed top-0 left-0 z-50 h-screen w-72 max-w-[85vw] bg-white border-r shadow-xl transition-transform duration-300 ${
            menuOpen ? "translate-x-0" : "-translate-x-full"
          }`}
          onClick={(e) => e.stopPropagation()}
        >
          <div className="flex items-center justify-between h-14 px-4 border-b">
            <span className="text-sm font-medium">メニュー</span>
            <button
              className="relative inline-flex items-center justify-center h-8 w-8 rounded hover:bg-gray-100 focus:outline-none focus-visible:ring-2 focus-visible:ring-gray-300"
              aria-label="閉じる"
              onClick={() => setMenuOpen(false)}
            >
              <span className="sr-only">閉じる</span>
              <span
                aria-hidden
                className="pointer-events-none absolute block h-0.5 w-4 bg-gray-800 rotate-45"
              />
              <span
                aria-hidden
                className="pointer-events-none absolute block h-0.5 w-4 bg-gray-800 -rotate-45"
              />
            </button>
          </div>
          <div className="flex flex-col py-2">
            <Link
              href="/about"
              className="px-4 py-2 text-sm hover:bg-gray-50"
              onClick={() => setMenuOpen(false)}
            >
              About
            </Link>
            <Link
              href="/contact"
              className="px-4 py-2 text-sm hover:bg-gray-50"
              onClick={() => setMenuOpen(false)}
            >
              Contact
            </Link>
            <Link
              href="/portfolio"
              className="px-4 py-2 text-sm hover:bg-gray-50"
              onClick={() => setMenuOpen(false)}
            >
              Portfolio
            </Link>
            <Link
              href="/archive"
              className="px-4 py-2 text-sm hover:bg-gray-50"
              onClick={() => setMenuOpen(false)}
            >
              Archive
            </Link>
            <Link
              href="/tech-stack"
              className="px-4 py-2 text-sm hover:bg-gray-50"
              onClick={() => setMenuOpen(false)}
            >
              Tech Stack
            </Link>
            <Link
              href="/privacy"
              className="px-4 py-2 text-sm hover:bg-gray-50"
              onClick={() => setMenuOpen(false)}
            >
              Privacy
            </Link>
            <Link
              href="/sitemap"
              className="px-4 py-2 text-sm hover:bg-gray-50"
              onClick={() => setMenuOpen(false)}
            >
              Sitemap
            </Link>
            <div className="border-t border-gray-200 mt-2 pt-2">
              <a
                href="https://github.com/tamaco489"
                target="_blank"
                rel="noopener noreferrer"
                className="px-4 py-2 text-sm hover:bg-gray-50 flex items-center"
              >
                <span>GitHub</span>
                <svg className="w-3 h-3 ml-1" fill="currentColor" viewBox="0 0 20 20">
                  <path fillRule="evenodd" d="M4.25 5.5a.75.75 0 00-.75.75v8.5c0 .414.336.75.75.75h8.5a.75.75 0 00.75-.75v-4a.75.75 0 011.5 0v4A2.25 2.25 0 0112.75 17h-8.5A2.25 2.25 0 012 14.75v-8.5A2.25 2.25 0 014.25 4h5a.75.75 0 010 1.5h-5z" clipRule="evenodd" />
                  <path fillRule="evenodd" d="M6.194 12.753a.75.75 0 001.06.053L16.5 4.44v2.81a.75.75 0 001.5 0v-4.5a.75.75 0 00-.75-.75h-4.5a.75.75 0 000 1.5h2.553l-9.056 8.194a.75.75 0 00-.053 1.06z" clipRule="evenodd" />
                </svg>
              </a>
            </div>
          </div>
        </nav>

        {/* 中央: ロゴ + タイトル */}
        <div className="justify-self-center">
          <Link
            href="/"
            className="flex items-center gap-3 text-3xl font-semibold tracking-wide text-white hover:text-gray-300 transition-colors"
          >
            <Image
              src="/blog/icon/favicon.png"
              alt="tamaco-blog ロゴ"
              width={40}
              height={40}
              className="rounded-lg"
              priority
            />
            tamaco-blog
          </Link>
        </div>

        {/* 右: アバターアイコン */}
        <div className="justify-self-end">
          <Link href="/profile" className="inline-flex items-center">
            <Image
              src="/profile/icon/avatar.png"
              alt="プロフィールアイコン"
              width={56}
              height={56}
              className="rounded-full border"
              priority
            />
          </Link>
        </div>
      </div>
    </header>
  );
}

export default Header;
