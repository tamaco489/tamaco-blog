package tag

import (
	"context"

	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/gen"
)

// GetArticlesByTagIdUseCase handles getting articles by tag ID
type GetArticlesByTagIdUseCase interface {
	Execute(ctx context.Context, tagId string, params gen.GetArticlesByTagIdParams) (*gen.ArticleList, error)
}

type getArticlesByTagIdUseCase struct{}

// NewGetArticlesByTagIdUseCase creates a new get articles by tag ID usecase
func NewGetArticlesByTagIdUseCase() GetArticlesByTagIdUseCase {
	return &getArticlesByTagIdUseCase{}
}

// Execute implements GetArticlesByTagIdUseCase
func (u *getArticlesByTagIdUseCase) Execute(ctx context.Context, tagId string, params gen.GetArticlesByTagIdParams) (*gen.ArticleList, error) {
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
