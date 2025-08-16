import { getMockArticleBySlug } from "@/lib/mock";
import { NextRequest, NextResponse } from "next/server";

export async function GET(
  request: NextRequest,
  { params }: { params: { slug: string } }
) {
  const { slug } = params;

  try {
    const article = getMockArticleBySlug(slug);

    if (!article) {
      return NextResponse.json(
        { error: "記事が見つかりません" },
        { status: 404 }
      );
    }

    return NextResponse.json(article, {
      status: 200,
      headers: {
        "Content-Type": "application/json",
      },
    });
  } catch (error) {
    console.error("記事データの取得に失敗しました:", error);
    return NextResponse.json(
      { error: "記事データの取得に失敗しました" },
      { status: 500 }
    );
  }
}
