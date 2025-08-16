import TechStack from "@/app/tech-stack/page";
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

describe("Tech Stack Page", () => {
  beforeEach(() => {
    jest.spyOn(console, "log").mockImplementation(() => {});
  });

  afterEach(() => {
    jest.restoreAllMocks();
  });

  test("ページが正しくレンダリングされる", () => {
    render(<TechStack />);

    expect(
      screen.getByRole("heading", { name: "Tech Stack" })
    ).toBeInTheDocument();
    expect(screen.getByText("使用技術・スキルセット")).toBeInTheDocument();
  });

  test("技術カテゴリが表示される", () => {
    render(<TechStack />);

    expect(screen.getByText("Frontend")).toBeInTheDocument();
    expect(screen.getByText("Backend")).toBeInTheDocument();
    expect(screen.getByText("Database")).toBeInTheDocument();
    expect(screen.getByText("Infrastructure")).toBeInTheDocument();
    expect(screen.getByText("DevOps")).toBeInTheDocument();
    expect(screen.getByText("Tools")).toBeInTheDocument();
  });

  test("フロントエンド技術が表示される", () => {
    render(<TechStack />);

    expect(screen.getByText("React")).toBeInTheDocument();
    expect(screen.getByText("Next.js")).toBeInTheDocument();
    expect(screen.getByText("TypeScript")).toBeInTheDocument();
    expect(screen.getAllByText("TailwindCSS").length).toBeGreaterThan(0);
  });

  test("バックエンド技術が表示される", () => {
    render(<TechStack />);

    expect(screen.getByText("Go")).toBeInTheDocument();
    expect(screen.getByText("Node.js")).toBeInTheDocument();
    expect(screen.getByText("Python")).toBeInTheDocument();
  });

  test("データベース技術が表示される", () => {
    render(<TechStack />);

    expect(screen.getAllByText("PostgreSQL").length).toBeGreaterThan(0);
    expect(screen.getByText("MySQL")).toBeInTheDocument();
    expect(screen.getByText("DynamoDB")).toBeInTheDocument();
  });

  test("インフラ技術が表示される", () => {
    render(<TechStack />);

    expect(screen.getAllByText("AWS").length).toBeGreaterThan(0);
    expect(screen.getAllByText("Docker").length).toBeGreaterThan(0);
    expect(screen.getByText("Terraform")).toBeInTheDocument();
  });

  test("習熟度が表示される", () => {
    render(<TechStack />);

    const activeItems = screen.getAllByText("使用中");
    const experiencedItems = screen.getAllByText("使用経験あり");
    const learningItems = screen.getAllByText("学習中");

    expect(activeItems.length).toBeGreaterThan(0);
    expect(experiencedItems.length).toBeGreaterThan(0);
    expect(learningItems.length).toBeGreaterThan(0);
  });

  test("技術カテゴリが適切に表示される", () => {
    render(<TechStack />);

    expect(screen.getByText("Frontend")).toBeInTheDocument();
    expect(screen.getByText("Backend")).toBeInTheDocument();
    expect(screen.getByText("Infrastructure")).toBeInTheDocument();
  });

  test("使用中の技術が表示される", () => {
    render(<TechStack />);

    expect(screen.getAllByText("使用中").length).toBeGreaterThan(0);
    expect(
      screen.getByText(/React ベースのフルスタックフレームワーク/)
    ).toBeInTheDocument();
  });


  test("技術の説明が表示される", () => {
    render(<TechStack />);

    expect(
      screen.getByText(/ユーザーインターフェース構築ライブラリ/)
    ).toBeInTheDocument();
    expect(
      screen.getByText(/React ベースのフルスタックフレームワーク/)
    ).toBeInTheDocument();
    expect(
      screen.getByText(/型安全性を提供する JavaScript のスーパーセット/)
    ).toBeInTheDocument();
  });

  test("技術の経験情報が表示される", () => {
    render(<TechStack />);

    expect(screen.getAllByText("使用中").length).toBeGreaterThan(0);
    expect(screen.getAllByText(/主な特徴・機能/).length).toBeGreaterThan(0);
  });
});
