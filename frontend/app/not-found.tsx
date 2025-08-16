"use client";

import Link from "next/link";
import { useEffect } from "react";

export default function NotFound() {
  useEffect(() => {
    console.log("🔍 ページが見つからないようです... tamaco-blog/404");
  }, []);

  return (
    <div className="min-h-[calc(100vh-7rem)] bg-gray-50 flex items-center justify-center px-4">
      <div className="max-w-lg mx-auto text-center">
        <div className="mb-8">
          <h1 className="text-9xl font-bold text-gray-300 mb-4">404</h1>
          <h2 className="text-2xl font-semibold text-gray-800 mb-4">
            ページが見つかりません
          </h2>
          <p className="text-gray-600 mb-8">
            お探しのページは存在しないか、移動された可能性があります。
            <br />
            <span className="text-sm text-gray-500 italic">
              console.log("存在しないページですね...🤔");
            </span>
          </p>
        </div>

        <div className="space-y-4">
          <Link
            href="/"
            className="inline-block bg-gray-900 hover:bg-gray-800 text-white font-medium py-3 px-6 rounded-lg transition-colors"
          >
            ホームに戻る
          </Link>
          
          <div className="text-sm text-gray-500">
            <p>または以下のページをお試しください：</p>
            <div className="flex flex-wrap justify-center gap-2 mt-2">
              <Link href="/about" className="text-blue-600 hover:underline">
                About
              </Link>
              <span>•</span>
              <Link href="/contact" className="text-blue-600 hover:underline">
                Contact
              </Link>
              <span>•</span>
              <Link href="/portfolio" className="text-blue-600 hover:underline">
                Portfolio
              </Link>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}