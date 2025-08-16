import About from "@/app/about/page";
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

describe("About Page", () => {
  beforeEach(() => {
    jest.spyOn(console, "log").mockImplementation(() => {});
  });

  afterEach(() => {
    jest.restoreAllMocks();
  });

  test("ページが正しくレンダリングされる", () => {
    render(<About />);

    expect(screen.getByRole("heading", { name: "About" })).toBeInTheDocument();
    expect(screen.getByText("このブログについて")).toBeInTheDocument();
  });

  test("ブログの説明が表示される", () => {
    render(<About />);

    expect(screen.getByText("tamaco-blog について")).toBeInTheDocument();
    expect(
      screen.getByText(/画面の向こうから、サーバーの奥まで/)
    ).toBeInTheDocument();
    expect(
      screen.getByText(/技術的な発見や学びを共有するテックブログです/)
    ).toBeInTheDocument();
  });

  test("トピックセクションが表示される", () => {
    render(<About />);

    expect(screen.getByText("このブログで扱うトピック")).toBeInTheDocument();
    expect(
      screen.getByText(/フロントエンド開発（React, Next.js, TypeScript）/)
    ).toBeInTheDocument();
    expect(
      screen.getByText(/バックエンド開発（Go, Node.js）/)
    ).toBeInTheDocument();
    expect(
      screen.getByText(/インフラ・クラウド（AWS, Docker）/)
    ).toBeInTheDocument();
  });

  test("技術スタックが表示される", () => {
    render(<About />);

    expect(screen.getByText("技術スタック")).toBeInTheDocument();
    expect(
      screen.getByText(/Next.js 14, TypeScript, TailwindCSS/)
    ).toBeInTheDocument();
    expect(screen.getByText(/Go \(予定\)/)).toBeInTheDocument();
    expect(
      screen.getByText(/Supabase PostgreSQL \(予定\)/)
    ).toBeInTheDocument();
  });


  test("学習プロジェクトとしての説明が表示される", () => {
    render(<About />);

    expect(
      screen.getByText(
        /このサイトはモダンな技術スタックを使用した学習・実験プロジェクトとしても機能しています/
      )
    ).toBeInTheDocument();
  });
});
