import Pagination from "@/components/Pagination";
import { fireEvent, render, screen } from "@testing-library/react";

const mockOnPageChange = jest.fn();

describe("Pagination", () => {
  beforeEach(() => {
    mockOnPageChange.mockClear();
  });

  test("renders pagination controls", () => {
    render(
      <Pagination
        currentPage={1}
        totalPages={5}
        onPageChange={mockOnPageChange}
      />
    );

    expect(screen.getByText("前へ")).toBeInTheDocument();
    expect(screen.getByText("次へ")).toBeInTheDocument();
    expect(screen.getByText("1")).toBeInTheDocument();
  });

  test("disables previous button on first page", () => {
    render(
      <Pagination
        currentPage={1}
        totalPages={5}
        onPageChange={mockOnPageChange}
      />
    );

    const prevButton = screen.getByText("前へ");
    expect(prevButton).toBeDisabled();
  });

  test("disables next button on last page", () => {
    render(
      <Pagination
        currentPage={5}
        totalPages={5}
        onPageChange={mockOnPageChange}
      />
    );

    const nextButton = screen.getByText("次へ");
    expect(nextButton).toBeDisabled();
  });

  test("calls onPageChange when page number is clicked", () => {
    render(
      <Pagination
        currentPage={2}
        totalPages={5}
        onPageChange={mockOnPageChange}
      />
    );

    const pageButton = screen.getByText("3");
    fireEvent.click(pageButton);

    expect(mockOnPageChange).toHaveBeenCalledWith(3);
  });

  test("calls onPageChange when next button is clicked", () => {
    render(
      <Pagination
        currentPage={2}
        totalPages={5}
        onPageChange={mockOnPageChange}
      />
    );

    const nextButton = screen.getByText("次へ");
    fireEvent.click(nextButton);

    expect(mockOnPageChange).toHaveBeenCalledWith(3);
  });

  test("calls onPageChange when previous button is clicked", () => {
    render(
      <Pagination
        currentPage={3}
        totalPages={5}
        onPageChange={mockOnPageChange}
      />
    );

    const prevButton = screen.getByText("前へ");
    fireEvent.click(prevButton);

    expect(mockOnPageChange).toHaveBeenCalledWith(2);
  });

  test("highlights current page", () => {
    render(
      <Pagination
        currentPage={3}
        totalPages={5}
        onPageChange={mockOnPageChange}
      />
    );

    const currentPageButton = screen.getByText("3");
    expect(currentPageButton).toHaveClass("bg-blue-600");
  });

  test("does not render when totalPages is 1", () => {
    const { container } = render(
      <Pagination
        currentPage={1}
        totalPages={1}
        onPageChange={mockOnPageChange}
      />
    );

    expect(container.firstChild).toBeNull();
  });

  test("shows ellipsis for large page counts", () => {
    render(
      <Pagination
        currentPage={5}
        totalPages={10}
        onPageChange={mockOnPageChange}
      />
    );

    const ellipsisElements = screen.getAllByText("...");
    expect(ellipsisElements.length).toBeGreaterThan(0);
  });
});
