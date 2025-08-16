import Portfolio from "@/app/portfolio/page";
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

describe("Portfolio Page", () => {
  beforeEach(() => {
    jest.spyOn(console, "log").mockImplementation(() => {});
  });

  afterEach(() => {
    jest.restoreAllMocks();
  });

  test("ページが正しくレンダリングされる", () => {
    render(<Portfolio />);

    expect(
      screen.getByRole("heading", { name: "Portfolio" })
    ).toBeInTheDocument();
    expect(screen.getByText("これまでの制作物・取り組み")).toBeInTheDocument();
  });

  test("プロジェクトが表示される", () => {
    render(<Portfolio />);

    expect(screen.getByText("tamaco-blog")).toBeInTheDocument();
    expect(
      screen.getByText(
        /Next.js \+ TypeScript \+ TailwindCSS で構築したテックブログ/
      )
    ).toBeInTheDocument();
    expect(screen.getByText("Go API Server")).toBeInTheDocument();
    expect(screen.getByText("AWS Infrastructure")).toBeInTheDocument();
  });

  test("技術タグが表示される", () => {
    render(<Portfolio />);

    expect(screen.getByText("Next.js 14")).toBeInTheDocument();
    expect(screen.getByText("TypeScript")).toBeInTheDocument();
    expect(screen.getByText("TailwindCSS")).toBeInTheDocument();
    expect(screen.getByText("Go")).toBeInTheDocument();
    expect(screen.getByText("AWS")).toBeInTheDocument();
  });

  test("プロジェクトステータスが表示される", () => {
    render(<Portfolio />);

    expect(screen.getByText("開発中")).toBeInTheDocument();
    expect(screen.getAllByText("企画中").length).toBeGreaterThan(0);
  });

  test("GitHub リンクが存在する", () => {
    render(<Portfolio />);

    const githubLinks = screen.getAllByRole("link", { name: /GitHub/i });
    expect(githubLinks.length).toBeGreaterThan(0);
  });


  test("プロジェクトの説明が表示される", () => {
    render(<Portfolio />);

    expect(
      screen.getByText(
        /Next\.js \+ TypeScript \+ TailwindCSS で構築したテックブログ/
      )
    ).toBeInTheDocument();
    expect(
      screen.getByText(/ブログのバックエンド API（予定）/)
    ).toBeInTheDocument();
    expect(
      screen.getByText(/クラウドインフラ構築（予定）/)
    ).toBeInTheDocument();
  });
});
