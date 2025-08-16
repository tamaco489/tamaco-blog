import ArticleDetailPage from "@/app/articles/[slug]/page";
import "@testing-library/jest-dom";
import { render, screen } from "@testing-library/react";

jest.mock("next/link", () => {
  const MockedLink = ({
    children,
    href,
  }: {
    children: React.ReactNode;
    href: string;
  }) => {
    return <a href={href}>{children}</a>;
  };
  MockedLink.displayName = "Link";
  return MockedLink;
});

jest.mock("next/navigation", () => ({
  notFound: jest.fn(),
}));

describe("ArticleDetailPage", () => {
  beforeEach(() => {
    jest.spyOn(console, "log").mockImplementation(() => {});
  });

  afterEach(() => {
    jest.restoreAllMocks();
  });

  test("既存の記事スラッグで正しくレンダリングされる", () => {
    const params = { slug: "nextjs-app-router-getting-started" };
    render(<ArticleDetailPage params={params} />);

    expect(
      screen.getByText("Next.js App Routerで始める最新のReact開発")
    ).toBeInTheDocument();
    expect(
      screen.getByText(/Next.js 13の新機能App Routerの基本的な使い方と/)
    ).toBeInTheDocument();
  });

  test("パンくずリストが表示される", () => {
    const params = { slug: "nextjs-app-router-getting-started" };
    render(<ArticleDetailPage params={params} />);

    expect(screen.getByRole("link", { name: "Home" })).toHaveAttribute(
      "href",
      "/"
    );
    expect(screen.getByText("記事詳細")).toBeInTheDocument();
  });

  test("記事のメタ情報が表示される", () => {
    const params = { slug: "nextjs-app-router-getting-started" };
    render(<ArticleDetailPage params={params} />);

    expect(screen.getByText("Next.js")).toBeInTheDocument(); // カテゴリ
    expect(screen.getByText("1,250 views")).toBeInTheDocument(); // 閲覧数
    expect(screen.getByText("by")).toBeInTheDocument(); // 著者情報
  });

  test("タグが表示される", () => {
    const params = { slug: "nextjs-app-router-getting-started" };
    render(<ArticleDetailPage params={params} />);

    expect(screen.getByText("#フロントエンド")).toBeInTheDocument();
    expect(screen.getByText("#JavaScript")).toBeInTheDocument();
  });

  test("記事本文が表示される", () => {
    const params = { slug: "nextjs-app-router-getting-started" };
    render(<ArticleDetailPage params={params} />);

    expect(
      screen.getByText(
        /Next.js 13で導入されたApp Routerについて詳しく解説します/
      )
    ).toBeInTheDocument();
  });

  test("記事一覧に戻るリンクが表示される", () => {
    const params = { slug: "nextjs-app-router-getting-started" };
    render(<ArticleDetailPage params={params} />);

    const backLink = screen.getByRole("link", { name: "記事一覧に戻る" });
    expect(backLink).toHaveAttribute("href", "/");
  });

  test("最終更新日が表示される", () => {
    const params = { slug: "nextjs-app-router-getting-started" };
    render(<ArticleDetailPage params={params} />);

    expect(screen.getByText("最終更新")).toBeInTheDocument();
  });

  test("存在しない記事スラッグの場合、記事データが取得できない", async () => {
    // 直接getMockArticleBySlugをテストして、nullが返されることを確認
    const { getMockArticleBySlug } = await import("@/lib/mock");
    const result = getMockArticleBySlug("non-existent-article");

    expect(result).toBeNull();
  });
});
