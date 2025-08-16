import ArticleCard from "@/components/ArticleCard";
import { mockArticles } from "@/lib/mock";
import { render, screen } from "@testing-library/react";

const mockArticle = mockArticles[0];

describe("ArticleCard", () => {
  test("renders article title", () => {
    render(<ArticleCard article={mockArticle} />);

    expect(screen.getByText(mockArticle.title)).toBeInTheDocument();
  });

  test("renders article excerpt", () => {
    render(<ArticleCard article={mockArticle} />);

    expect(screen.getByText(mockArticle.excerpt!)).toBeInTheDocument();
  });

  test("renders category with correct color", () => {
    render(<ArticleCard article={mockArticle} />);

    const categoryElement = screen.getByText(mockArticle.category!.name);
    expect(categoryElement).toBeInTheDocument();
  });

  test("renders view count", () => {
    render(<ArticleCard article={mockArticle} />);

    const viewCount = `${mockArticle.view_count.toLocaleString()} views`;
    expect(screen.getByText(viewCount)).toBeInTheDocument();
  });

  test("renders published date", () => {
    render(<ArticleCard article={mockArticle} />);

    const expectedDate = new Date(mockArticle.published_at!).toLocaleDateString(
      "ja-JP"
    );
    expect(screen.getByText(expectedDate)).toBeInTheDocument();
  });

  test("renders tags when available", () => {
    render(<ArticleCard article={mockArticle} />);

    if (mockArticle.tags && mockArticle.tags.length > 0) {
      const firstTag = mockArticle.tags[0];
      expect(screen.getByText(`#${firstTag.name}`)).toBeInTheDocument();
    }
  });

  test("handles article without category gracefully", () => {
    const articleWithoutCategory = {
      ...mockArticle,
      category: undefined,
      category_id: undefined,
    };

    render(<ArticleCard article={articleWithoutCategory} />);

    expect(screen.getByText(mockArticle.title)).toBeInTheDocument();
  });

  test("handles article without tags gracefully", () => {
    const articleWithoutTags = {
      ...mockArticle,
      tags: undefined,
    };

    render(<ArticleCard article={articleWithoutTags} />);

    expect(screen.getByText(mockArticle.title)).toBeInTheDocument();
  });
});
