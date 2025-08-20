package category

import (
	"context"

	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/gen"
)

// GetArticlesByCategoryIdUseCase handles getting articles by category ID
type GetArticlesByCategoryIdUseCase interface {
	Execute(ctx context.Context, categoryId string, params gen.GetArticlesByCategoryIdParams) (*gen.ArticleList, error)
}

type getArticlesByCategoryIdUseCase struct{}

// NewGetArticlesByCategoryIdUseCase creates a new get articles by category ID usecase
func NewGetArticlesByCategoryIdUseCase() GetArticlesByCategoryIdUseCase {
	return &getArticlesByCategoryIdUseCase{}
}

// Execute implements GetArticlesByCategoryIdUseCase
func (u *getArticlesByCategoryIdUseCase) Execute(ctx context.Context, categoryId string, params gen.GetArticlesByCategoryIdParams) (*gen.ArticleList, error) {
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
