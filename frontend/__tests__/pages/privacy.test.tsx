import Privacy from "@/app/privacy/page";
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

describe("Privacy Page", () => {
  beforeEach(() => {
    jest.spyOn(console, "log").mockImplementation(() => {});
  });

  afterEach(() => {
    jest.restoreAllMocks();
  });

  test("ページが正しくレンダリングされる", () => {
    render(<Privacy />);

    expect(screen.getByRole("heading", { name: "プライバシーポリシー" })).toBeInTheDocument();
    expect(screen.getByText("アクセス解析とCookieの使用について")).toBeInTheDocument();
  });

  test("アクセス解析ツールセクションが表示される", () => {
    render(<Privacy />);

    expect(screen.getByText("アクセス解析ツールについて")).toBeInTheDocument();
    expect(
      screen.getByText(/当サイトでは、サイトの利用状況を把握し、より良いサービスを提供するために以下のツールを使用しています/)
    ).toBeInTheDocument();
    expect(screen.getByText(/Google Analytics/)).toBeInTheDocument();
    expect(screen.getByText(/New Relic/)).toBeInTheDocument();
  });

  test("Cookieの使用セクションが表示される", () => {
    render(<Privacy />);

    expect(screen.getByText("Cookieの使用について")).toBeInTheDocument();
    expect(
      screen.getByText(/当サイトでは、アクセス解析やパフォーマンス向上のためにCookieを使用しています/)
    ).toBeInTheDocument();
    expect(screen.getByText(/サイトの利用状況の分析/)).toBeInTheDocument();
    expect(screen.getByText(/ユーザーエクスペリエンスの向上/)).toBeInTheDocument();
    expect(screen.getByText(/サイトのパフォーマンス測定/)).toBeInTheDocument();
  });

  test("お問い合わせセクションが表示される", () => {
    render(<Privacy />);

    expect(screen.getByText("お問い合わせについて")).toBeInTheDocument();
    expect(
      screen.getByText(/お問い合わせフォームをご利用いただく際は、お名前とメールアドレスをお預かりします/)
    ).toBeInTheDocument();
  });

  test("ポリシーの変更セクションが表示される", () => {
    render(<Privacy />);

    expect(screen.getByText("ポリシーの変更")).toBeInTheDocument();
    expect(
      screen.getByText(/このプライバシーポリシーは、必要に応じて変更される場合があります/)
    ).toBeInTheDocument();
  });

  test("要約セクションが表示される", () => {
    render(<Privacy />);

    expect(screen.getByText("📋 要約")).toBeInTheDocument();
    expect(
      screen.getByText(/当サイトは技術ブログとして運営されており、閲覧に際して個人情報の登録は不要です/)
    ).toBeInTheDocument();
    expect(
      screen.getByText(/アクセス解析のためにCookieを使用していますが、これらは匿名化されており/)
    ).toBeInTheDocument();
  });

  test("最終更新日と運営者情報が表示される", () => {
    render(<Privacy />);

    expect(screen.getByText(/最終更新日: 2024年8月16日/)).toBeInTheDocument();
    expect(screen.getByText(/運営者: tamaco489/)).toBeInTheDocument();
    expect(
      screen.getByText(/このポリシーに関するご質問は、お問い合わせページよりご連絡ください/)
    ).toBeInTheDocument();
  });
});