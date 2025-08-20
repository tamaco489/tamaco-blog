// Package article provides article-related business logic.
package article

import (
	"context"
	"log/slog"

	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/gen"
)

// GetArticlesUseCase handles getting articles list
type GetArticlesUseCase interface {
	GetArticles(ctx context.Context, params gen.GetArticlesParams) (*gen.ArticleList, error)
}

type getArticlesUseCase struct {
	// repository層は未実装のため、今回は空
}

// NewGetArticlesUseCase creates a new get articles usecase
func NewGetArticlesUseCase() GetArticlesUseCase {
	return &getArticlesUseCase{}
}

// GetArticles implements GetArticlesUseCase
func (u *getArticlesUseCase) GetArticles(ctx context.Context, params gen.GetArticlesParams) (*gen.ArticleList, error) {
	slog.InfoContext(ctx, "[TEST] GetArticles called", "params", params)

	// TODO: 実装予定
	return &gen.ArticleList{
		Articles: []gen.Article{},
		Pagination: struct {
			CurrentPage int   `json:"current_page"`
			HasNext     *bool `json:"has_next,omitempty"`
			HasPrev     *bool `json:"has_prev,omitempty"`
			PageSize    int   `json:"page_size"`
			TotalCount  int   `json:"total_count"`
			TotalPages  int   `json:"total_pages"`
		}{
			CurrentPage: 1,
			PageSize:    20,
			TotalCount:  0,
			TotalPages:  0,
		},
	}, nil
}
