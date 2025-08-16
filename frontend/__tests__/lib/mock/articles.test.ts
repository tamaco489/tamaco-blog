import { getMockArticles, mockArticles } from "@/lib/mock/articles";

describe("getMockArticles", () => {
  test("should return first page with default limit", () => {
    const result = getMockArticles();

    expect(result.data).toHaveLength(10);
    expect(result.pagination.page).toBe(1);
    expect(result.pagination.limit).toBe(10);
    expect(result.pagination.total).toBe(mockArticles.length);
    expect(result.pagination.totalPages).toBe(
      Math.ceil(mockArticles.length / 10)
    );
    expect(result.pagination.hasNext).toBe(true);
    expect(result.pagination.hasPrev).toBe(false);
  });

  test("should return second page correctly", () => {
    const result = getMockArticles(2, 10);

    expect(result.data).toHaveLength(5); // 残り5件
    expect(result.pagination.page).toBe(2);
    expect(result.pagination.hasNext).toBe(false);
    expect(result.pagination.hasPrev).toBe(true);
  });

  test("should handle custom limit", () => {
    const result = getMockArticles(1, 5);

    expect(result.data).toHaveLength(5);
    expect(result.pagination.limit).toBe(5);
    expect(result.pagination.totalPages).toBe(3);
  });

  test("should return empty array for out of range page", () => {
    const result = getMockArticles(10, 10);

    expect(result.data).toHaveLength(0);
    expect(result.pagination.hasNext).toBe(false);
    expect(result.pagination.hasPrev).toBe(true);
  });

  test("should have correct article structure", () => {
    const result = getMockArticles(1, 1);
    const article = result.data[0];

    expect(article).toHaveProperty("id");
    expect(article).toHaveProperty("title");
    expect(article).toHaveProperty("slug");
    expect(article).toHaveProperty("content");
    expect(article).toHaveProperty("category");
    expect(article).toHaveProperty("author");
    expect(article).toHaveProperty("tags");
    expect(Array.isArray(article.tags)).toBe(true);
  });
});
