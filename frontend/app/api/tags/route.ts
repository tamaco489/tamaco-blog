import { mockTags } from "@/lib/mock";
import { NextResponse } from "next/server";

export async function GET() {
  try {
    return NextResponse.json(mockTags, {
      status: 200,
      headers: {
        "Content-Type": "application/json",
      },
    });
  } catch (error) {
    console.error("タグデータの取得に失敗しました:", error);
    return NextResponse.json(
      { error: "タグデータの取得に失敗しました" },
      { status: 500 }
    );
  }
}
