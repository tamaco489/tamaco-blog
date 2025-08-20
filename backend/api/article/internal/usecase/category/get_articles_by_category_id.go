package category

import (
	"context"

	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/gen"
)

type GetArticlesByCategoryIDUseCase interface {
	Execute(ctx context.Context, categoryID string, params gen.GetArticlesByCategoryIDParams) (*gen.ArticleList, error)
}

type getArticlesByCategoryIDUseCase struct{}

func NewGetArticlesByCategoryIDUseCase() GetArticlesByCategoryIDUseCase {
	return &getArticlesByCategoryIDUseCase{}
}

func (u *getArticlesByCategoryIDUseCase) Execute(ctx context.Context, categoryID string, params gen.GetArticlesByCategoryIDParams) (*gen.ArticleList, error) {
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
