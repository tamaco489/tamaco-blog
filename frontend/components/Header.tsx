"use client";
import Link from "next/link";
import Image from "next/image";
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
    <header className="w-full border-b">
      <div className="w-full h-28 grid grid-cols-3 items-center px-4">
        {/* 左: ハンバーガーメニュー */}
        <div ref={menuRef} className="relative justify-self-start">
          <button
            className="inline-flex flex-col items-center justify-center gap-1.5 h-10 w-10 rounded-md hover:bg-gray-100 focus:outline-none focus-visible:ring-2 focus-visible:ring-gray-300"
            aria-haspopup="true"
            aria-expanded={menuOpen}
            aria-label="メニュー"
            onClick={() => setMenuOpen((v) => !v)}
          >
            <span className="sr-only">メニュー</span>
            <span className="block h-0.5 w-6 bg-gray-800" />
            <span className="block h-0.5 w-6 bg-gray-800" />
            <span className="block h-0.5 w-6 bg-gray-800" />
          </button>
        </div>

        {/* スライドメニュー（左から、全高） */}
        <div
          className={`fixed inset-0 z-40 transition-opacity duration-300 bg-black/40 ${menuOpen ? "opacity-100 pointer-events-auto" : "opacity-0 pointer-events-none"}`}
          onClick={() => setMenuOpen(false)}
          aria-hidden={!menuOpen}
        />

        <nav
          role="dialog"
          aria-modal={menuOpen}
          aria-label="メニュー"
          className={`fixed top-0 left-0 z-50 h-screen w-72 max-w-[85vw] bg-white border-r shadow-xl transition-transform duration-300 ${menuOpen ? "translate-x-0" : "-translate-x-full"}`}
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
              <span aria-hidden className="pointer-events-none absolute block h-0.5 w-4 bg-gray-800 rotate-45" />
              <span aria-hidden className="pointer-events-none absolute block h-0.5 w-4 bg-gray-800 -rotate-45" />
            </button>
          </div>
          <div className="flex flex-col py-2">
            <Link href="/profile" className="px-4 py-2 text-sm hover:bg-gray-50">
              プロフィール
            </Link>
            <a
              href="https://github.com/tamaco489/tamaco-blog"
              target="_blank"
              rel="noopener noreferrer"
              className="px-4 py-2 text-sm hover:bg-gray-50"
            >
              GitHub
            </a>
          </div>
        </nav>

        {/* 中央: タイトル */}
        <div className="justify-self-center">
          <Link href="/" className="text-3xl font-semibold tracking-wide">
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
