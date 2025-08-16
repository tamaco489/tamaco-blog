import NotFound from "@/app/not-found";
import "@testing-library/jest-dom";
import { render, screen } from "@testing-library/react";

// Next.js Link をモック
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

describe("NotFound Page", () => {
  beforeEach(() => {
    // console.log をモック
    jest.spyOn(console, "log").mockImplementation(() => {});
  });

  afterEach(() => {
    jest.restoreAllMocks();
  });

  test("404ページが正しくレンダリングされる", () => {
    render(<NotFound />);

    expect(screen.getByText("404")).toBeInTheDocument();
    expect(screen.getByText("ページが見つかりません")).toBeInTheDocument();
    expect(screen.getByText("ホームに戻る")).toBeInTheDocument();
  });

  test("説明文が表示される", () => {
    render(<NotFound />);

    expect(
      screen.getByText(
        "お探しのページは存在しないか、移動された可能性があります。"
      )
    ).toBeInTheDocument();
  });

  test("ナビゲーションリンクが正しく表示される", () => {
    render(<NotFound />);

    const homeLink = screen.getByRole("link", { name: "ホームに戻る" });
    expect(homeLink).toHaveAttribute("href", "/");

    const aboutLink = screen.getByRole("link", { name: "About" });
    expect(aboutLink).toHaveAttribute("href", "/about");

    const contactLink = screen.getByRole("link", { name: "Contact" });
    expect(contactLink).toHaveAttribute("href", "/contact");

    const portfolioLink = screen.getByRole("link", { name: "Portfolio" });
    expect(portfolioLink).toHaveAttribute("href", "/portfolio");
  });


  test("ホームボタンが表示される", () => {
    render(<NotFound />);

    const homeButton = screen.getByRole("link", { name: "ホームに戻る" });
    expect(homeButton).toBeInTheDocument();
    expect(homeButton).toHaveAttribute("href", "/");
  });
});
