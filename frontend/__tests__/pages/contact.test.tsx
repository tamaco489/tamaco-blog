import Contact from "@/app/contact/page";
import "@testing-library/jest-dom";
import { fireEvent, render, screen } from "@testing-library/react";

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

Object.defineProperty(window, "alert", {
  writable: true,
  value: jest.fn(),
});

describe("Contact Page", () => {
  beforeEach(() => {
    jest.spyOn(console, "log").mockImplementation(() => {});
    (window.alert as jest.Mock).mockClear();
  });

  afterEach(() => {
    jest.restoreAllMocks();
  });

  test("ページが正しくレンダリングされる", () => {
    render(<Contact />);

    expect(
      screen.getByRole("heading", { name: "Contact" })
    ).toBeInTheDocument();
    expect(
      screen.getByText("お気軽にお問い合わせください")
    ).toBeInTheDocument();
  });

  test("フォームフィールドが表示される", () => {
    render(<Contact />);

    expect(screen.getByLabelText("お名前 *")).toBeInTheDocument();
    expect(screen.getByLabelText("メールアドレス *")).toBeInTheDocument();
    expect(screen.getByLabelText("件名 *")).toBeInTheDocument();
    expect(screen.getByLabelText("メッセージ *")).toBeInTheDocument();
  });

  test("フォーム入力が正しく処理される", () => {
    render(<Contact />);

    const nameInput = screen.getByLabelText("お名前 *");
    const emailInput = screen.getByLabelText("メールアドレス *");
    const subjectInput = screen.getByLabelText("件名 *");
    const messageInput = screen.getByLabelText("メッセージ *");

    fireEvent.change(nameInput, { target: { value: "テスト太郎" } });
    fireEvent.change(emailInput, { target: { value: "test@example.com" } });
    fireEvent.change(subjectInput, { target: { value: "テスト件名" } });
    fireEvent.change(messageInput, { target: { value: "テストメッセージ" } });

    expect(nameInput).toHaveValue("テスト太郎");
    expect(emailInput).toHaveValue("test@example.com");
    expect(subjectInput).toHaveValue("テスト件名");
    expect(messageInput).toHaveValue("テストメッセージ");
  });

  test("フォーム送信時にアラートが表示される", () => {
    render(<Contact />);

    const submitButton = screen.getByRole("button", {
      name: "メッセージを送信",
    });
    const nameInput = screen.getByLabelText("お名前 *");
    const emailInput = screen.getByLabelText("メールアドレス *");
    const subjectInput = screen.getByLabelText("件名 *");
    const messageInput = screen.getByLabelText("メッセージ *");

    fireEvent.change(nameInput, { target: { value: "テスト太郎" } });
    fireEvent.change(emailInput, { target: { value: "test@example.com" } });
    fireEvent.change(subjectInput, { target: { value: "テスト件名" } });
    fireEvent.change(messageInput, { target: { value: "テストメッセージ" } });

    fireEvent.click(submitButton);

    expect(window.alert).toHaveBeenCalledWith(
      "現在は開発中のため、実際の送信はできません。"
    );
  });

  test("GitHub リンクが正しく設定される", () => {
    render(<Contact />);

    const githubLink = screen.getByRole("link", { name: "@tamaco489" });
    expect(githubLink).toHaveAttribute("href", "https://github.com/tamaco489");
    expect(githubLink).toHaveAttribute("target", "_blank");
    expect(githubLink).toHaveAttribute("rel", "noopener noreferrer");
  });

  test("開発中の注意メッセージが表示される", () => {
    render(<Contact />);

    expect(
      screen.getByText(/現在フォーム送信機能は実装されていません/)
    ).toBeInTheDocument();
  });

  test("よくあるお問い合わせセクションが表示される", () => {
    render(<Contact />);

    expect(screen.getByText("よくあるお問い合わせ")).toBeInTheDocument();
    expect(screen.getByText("技術的な質問")).toBeInTheDocument();
    expect(screen.getByText("コラボレーション")).toBeInTheDocument();
  });

});
