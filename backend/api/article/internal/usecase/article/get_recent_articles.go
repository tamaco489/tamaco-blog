package article

import (
	"context"

	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/gen"
)

// GetRecentArticlesUseCase handles getting recent articles
type GetRecentArticlesUseCase interface {
	GetRecentArticles(ctx context.Context, params gen.GetRecentArticlesParams) (*gen.ArticleList, error)
}

type getRecentArticlesUseCase struct{}

// NewGetRecentArticlesUseCase creates a new get recent articles usecase
func NewGetRecentArticlesUseCase() GetRecentArticlesUseCase {
	return &getRecentArticlesUseCase{}
}

// GetRecentArticles implements GetRecentArticlesUseCase
func (u *getRecentArticlesUseCase) GetRecentArticles(ctx context.Context, params gen.GetRecentArticlesParams) (*gen.ArticleList, error) {
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
