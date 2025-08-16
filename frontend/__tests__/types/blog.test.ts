import { mockArticles, mockCategories, mockTags, mockUser } from "@/lib/mock";
import { Article, Category, Tag, User } from "@/types/blog";

describe("Blog Types", () => {
  describe("Category type", () => {
    test("mock categories match Category interface", () => {
      mockCategories.forEach((category: Category) => {
        expect(category).toHaveProperty("id");
        expect(category).toHaveProperty("name");
        expect(category).toHaveProperty("slug");
        expect(category).toHaveProperty("display_order");
        expect(category).toHaveProperty("created_at");
        expect(category).toHaveProperty("updated_at");

        expect(typeof category.id).toBe("string");
        expect(typeof category.name).toBe("string");
        expect(typeof category.slug).toBe("string");
        expect(typeof category.display_order).toBe("number");
      });
    });
  });

  describe("Tag type", () => {
    test("mock tags match Tag interface", () => {
      mockTags.forEach((tag: Tag) => {
        expect(tag).toHaveProperty("id");
        expect(tag).toHaveProperty("name");
        expect(tag).toHaveProperty("slug");
        expect(tag).toHaveProperty("created_at");

        expect(typeof tag.id).toBe("string");
        expect(typeof tag.name).toBe("string");
        expect(typeof tag.slug).toBe("string");
      });
    });
  });

  describe("User type", () => {
    test("mock user matches User interface", () => {
      const user: User = mockUser;

      expect(user).toHaveProperty("id");
      expect(user).toHaveProperty("email");
      expect(user).toHaveProperty("name");
      expect(user).toHaveProperty("role");
      expect(user).toHaveProperty("created_at");
      expect(user).toHaveProperty("updated_at");

      expect(typeof user.id).toBe("string");
      expect(typeof user.email).toBe("string");
      expect(typeof user.name).toBe("string");
      expect(["admin", "author"]).toContain(user.role);
    });
  });

  describe("Article type", () => {
    test("mock articles match Article interface", () => {
      mockArticles.forEach((article: Article) => {
        expect(article).toHaveProperty("id");
        expect(article).toHaveProperty("title");
        expect(article).toHaveProperty("slug");
        expect(article).toHaveProperty("content");
        expect(article).toHaveProperty("status");
        expect(article).toHaveProperty("author_id");
        expect(article).toHaveProperty("view_count");
        expect(article).toHaveProperty("created_at");
        expect(article).toHaveProperty("updated_at");

        expect(typeof article.id).toBe("string");
        expect(typeof article.title).toBe("string");
        expect(typeof article.slug).toBe("string");
        expect(typeof article.content).toBe("string");
        expect(["draft", "published", "private"]).toContain(article.status);
        expect(typeof article.view_count).toBe("number");

        // Relations
        if (article.category) {
          expect(article.category).toHaveProperty("id");
          expect(article.category).toHaveProperty("name");
        }

        if (article.author) {
          expect(article.author).toHaveProperty("id");
          expect(article.author).toHaveProperty("name");
        }

        if (article.tags) {
          expect(Array.isArray(article.tags)).toBe(true);
          article.tags.forEach((tag) => {
            expect(tag).toHaveProperty("id");
            expect(tag).toHaveProperty("name");
          });
        }
      });
    });
  });
});
