import { mockCategories } from "@/lib/mock";
import { NextResponse } from "next/server";

export async function GET() {
  try {
    return NextResponse.json(mockCategories, {
      status: 200,
      headers: {
        "Content-Type": "application/json",
      },
    });
  } catch (error) {
    console.error("カテゴリデータの取得に失敗しました:", error);
    return NextResponse.json(
      { error: "カテゴリデータの取得に失敗しました" },
      { status: 500 }
    );
  }
}
