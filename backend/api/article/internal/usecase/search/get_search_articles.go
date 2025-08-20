package search

import (
	"context"

	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/gen"
)

// SearchArticlesUseCase handles searching articles
type SearchArticlesUseCase interface {
	SearchArticles(ctx context.Context, params gen.SearchArticlesParams) (*gen.ArticleList, error)
}

type searchArticlesUseCase struct{}

// NewSearchArticlesUseCase creates a new search articles usecase
func NewSearchArticlesUseCase() SearchArticlesUseCase {
	return &searchArticlesUseCase{}
}

// SearchArticles implements SearchArticlesUseCase
func (u *searchArticlesUseCase) SearchArticles(ctx context.Context, params gen.SearchArticlesParams) (*gen.ArticleList, error) {
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
