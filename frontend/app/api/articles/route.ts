import { getMockArticles } from "@/lib/mock";
import { NextRequest, NextResponse } from "next/server";

export async function GET(request: NextRequest) {
  const searchParams = request.nextUrl.searchParams;
  const page = parseInt(searchParams.get("page") || "1", 10);
  const limit = parseInt(searchParams.get("limit") || "10", 10);

  try {
    const articleData = getMockArticles(page, limit);

    return NextResponse.json(articleData, {
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
