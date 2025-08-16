import { Article } from "@/types/blog";
import Link from "next/link";

interface ArticleCardProps {
  article: Article;
}

export default function ArticleCard({ article }: ArticleCardProps) {
  return (
    <Link href={`/articles/${article.slug}`} className="block h-full">
      <div className="w-full bg-white border border-gray-200 rounded-lg shadow-md hover:shadow-lg transition-shadow duration-300 h-full flex flex-col cursor-pointer">
        <div className="p-6 flex flex-col flex-1">
          <h2 className="text-xl font-bold text-gray-900 mb-3 line-clamp-2 hover:text-blue-600 transition-colors break-words">
            {article.title}
          </h2>

          <p className="text-gray-600 text-sm mb-4 line-clamp-3 flex-1 break-words">
            {article.excerpt}
          </p>

          <div className="flex items-center justify-between text-xs text-gray-500 mb-4">
            <div className="flex items-center space-x-4">
              {article.category && (
                <span
                  className="px-2 py-1 rounded-full text-white text-xs font-medium"
                  style={{
                    backgroundColor: article.category.color || "#6B7280",
                  }}
                >
                  {article.category.name}
                </span>
              )}
              <span>
                {new Date(
                  article.published_at || article.created_at
                ).toLocaleDateString("ja-JP")}
              </span>
            </div>
            <span>{article.view_count.toLocaleString()} views</span>
          </div>

          {article.tags && article.tags.length > 0 && (
            <div className="flex flex-wrap gap-2">
              {article.tags.slice(0, 3).map((tag) => (
                <span
                  key={tag.id}
                  className="px-2 py-1 bg-gray-100 text-gray-700 rounded text-xs"
                >
                  #{tag.name}
                </span>
              ))}
            </div>
          )}
        </div>
      </div>
    </Link>
  );
}
