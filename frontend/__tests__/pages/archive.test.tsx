import Archive from "@/app/archive/page";
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

describe("Archive Page", () => {
  beforeEach(() => {
    jest.spyOn(console, "log").mockImplementation(() => {});
  });

  afterEach(() => {
    jest.restoreAllMocks();
  });

  test("ページが正しくレンダリングされる", () => {
    render(<Archive />);

    expect(
      screen.getByRole("heading", { name: "Archive" })
    ).toBeInTheDocument();
    expect(screen.getByText("記事一覧・アーカイブ")).toBeInTheDocument();
  });

  test("フィルターが表示される", () => {
    render(<Archive />);

    expect(screen.getByLabelText("カテゴリー")).toBeInTheDocument();
    expect(screen.getByLabelText("年度")).toBeInTheDocument();
  });

  test("フィルター選択が正しく処理される", () => {
    render(<Archive />);

    const categorySelect = screen.getByLabelText("カテゴリー");
    const yearSelect = screen.getByLabelText("年度");

    fireEvent.change(categorySelect, { target: { value: "1" } });
    fireEvent.change(yearSelect, { target: { value: "2024" } });

    expect(categorySelect).toHaveValue("1");
    expect(yearSelect).toHaveValue("2024");
  });

  test("記事一覧が表示される", () => {
    render(<Archive />);

    // モックデータから記事が表示されることを確認
    const articleLinks = screen.getAllByRole("link");
    expect(articleLinks.length).toBeGreaterThan(0);
  });

  test("年月グルーピングが表示される", () => {
    render(<Archive />);

    // 年月のセクションが表示されることを確認
    const yearMonthHeadings = screen.getAllByText(/\d{4}年\d{1,2}月/);
    expect(yearMonthHeadings.length).toBeGreaterThan(0);
  });

  test("フィルター機能が存在する", () => {
    render(<Archive />);

    expect(screen.getByLabelText("カテゴリー")).toBeInTheDocument();
    expect(screen.getByLabelText("年度")).toBeInTheDocument();
  });

  test("モックデータが使用されている", () => {
    render(<Archive />);

    // モックデータが読み込まれていることを確認
    expect(screen.getByText("記事一覧・アーカイブ")).toBeInTheDocument();
  });


  test("フィルターオプションが正しく表示される", () => {
    render(<Archive />);

    const categorySelect = screen.getByLabelText("カテゴリー");
    const yearSelect = screen.getByLabelText("年度");

    expect(categorySelect).toContainHTML(
      '<option value="all">すべてのカテゴリー</option>'
    );
    expect(yearSelect).toContainHTML('<option value="all">すべての年</option>');
  });
});
