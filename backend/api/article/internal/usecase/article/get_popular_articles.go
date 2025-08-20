package article

import (
	"context"

	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/gen"
)

// GetPopularArticlesUseCase handles getting popular articles
type GetPopularArticlesUseCase interface {
	GetPopularArticles(ctx context.Context, params gen.GetPopularArticlesParams) (*gen.ArticleList, error)
}

type getPopularArticlesUseCase struct{}

// NewGetPopularArticlesUseCase creates a new get popular articles usecase
func NewGetPopularArticlesUseCase() GetPopularArticlesUseCase {
	return &getPopularArticlesUseCase{}
}

// GetPopularArticles implements GetPopularArticlesUseCase
func (u *getPopularArticlesUseCase) GetPopularArticles(ctx context.Context, params gen.GetPopularArticlesParams) (*gen.ArticleList, error) {
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
