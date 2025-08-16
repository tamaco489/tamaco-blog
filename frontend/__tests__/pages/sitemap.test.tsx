import Sitemap from "@/app/sitemap/page";
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

describe("Sitemap Page", () => {
  beforeEach(() => {
    jest.spyOn(console, "log").mockImplementation(() => {});
  });

  afterEach(() => {
    jest.restoreAllMocks();
  });

  test("ページが正しくレンダリングされる", () => {
    render(<Sitemap />);

    expect(screen.getByRole("heading", { level: 1, name: "サイトマップ" })).toBeInTheDocument();
    expect(screen.getByText("当サイトの全ページ一覧")).toBeInTheDocument();
  });

  test("メインページセクションが表示される", () => {
    render(<Sitemap />);

    expect(screen.getByText("メインページ")).toBeInTheDocument();
    expect(screen.getByText("ホーム")).toBeInTheDocument();
    expect(screen.getByText("About")).toBeInTheDocument();
    expect(screen.getByText("お問い合わせ")).toBeInTheDocument();
    expect(screen.getByText("ポートフォリオ")).toBeInTheDocument();
    expect(screen.getByText("技術スタック")).toBeInTheDocument();
  });

  test("記事・コンテンツセクションが表示される", () => {
    render(<Sitemap />);

    expect(screen.getByText("記事・コンテンツ")).toBeInTheDocument();
    expect(screen.getByText("アーカイブ")).toBeInTheDocument();
    expect(screen.getByText("記事アーカイブ（今後実装予定）")).toBeInTheDocument();
  });

  test("法的ページセクションが表示される", () => {
    render(<Sitemap />);

    expect(screen.getByText("法的ページ")).toBeInTheDocument();
    expect(screen.getByText("プライバシーポリシー")).toBeInTheDocument();
    expect(screen.getByText("個人情報の取り扱い")).toBeInTheDocument();
    expect(screen.getByText("このページ")).toBeInTheDocument();
  });

  test("利用規約が表示されない", () => {
    render(<Sitemap />);

    expect(screen.queryByText("利用規約")).not.toBeInTheDocument();
    expect(screen.queryByText("サイト利用規約")).not.toBeInTheDocument();
  });

  test("適切なリンクが設定されている", () => {
    render(<Sitemap />);

    const links = screen.getAllByRole("link");
    const homeLink = links.find(link => link.getAttribute("href") === "/");
    const aboutLink = links.find(link => link.getAttribute("href") === "/about" && link.textContent?.includes("About"));
    const contactLink = links.find(link => link.getAttribute("href") === "/contact");
    const privacyLink = links.find(link => link.getAttribute("href") === "/privacy");

    expect(homeLink).toBeInTheDocument();
    expect(aboutLink).toBeInTheDocument();
    expect(contactLink).toBeInTheDocument();
    expect(privacyLink).toBeInTheDocument();
  });

  test("URLが正しく表示される", () => {
    render(<Sitemap />);

    expect(screen.getByText("/")).toBeInTheDocument();
    expect(screen.getByText("/about")).toBeInTheDocument();
    expect(screen.getByText("/contact")).toBeInTheDocument();
    expect(screen.getByText("/privacy")).toBeInTheDocument();
    expect(screen.getByText("/sitemap")).toBeInTheDocument();
  });

  test("今後実装予定のページセクションが表示される", () => {
    render(<Sitemap />);

    expect(screen.getByText("今後実装予定のページ")).toBeInTheDocument();
    expect(screen.getByText("記事一覧・詳細ページ（/articles）")).toBeInTheDocument();
    expect(screen.getByText("カテゴリ別記事ページ（/categories）")).toBeInTheDocument();
    expect(screen.getByText("タグ別記事ページ（/tags）")).toBeInTheDocument();
    expect(screen.getByText("検索機能（/search）")).toBeInTheDocument();
    expect(screen.getByText("管理画面（/admin）")).toBeInTheDocument();
  });

  test("開発状況セクションが表示される", () => {
    render(<Sitemap />);

    expect(screen.getByText("🛠️ 開発状況について")).toBeInTheDocument();
    expect(
      screen.getByText(/tamaco-blogは現在開発中のプロジェクトです/)
    ).toBeInTheDocument();
    expect(
      screen.getByText(/Next.js 14 \+ TypeScript \+ TailwindCSSを使用して構築されており/)
    ).toBeInTheDocument();
    
    const aboutPageLink = screen.getByRole("link", { name: "Aboutページ" });
    expect(aboutPageLink).toHaveAttribute("href", "/about");
  });
});