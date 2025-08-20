package tag

import (
	"context"

	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/gen"
)

type GetArticlesByTagIDUseCase interface {
	Execute(ctx context.Context, tagID string, params gen.GetArticlesByTagIDParams) (*gen.ArticleList, error)
}

type getArticlesByTagIDUseCase struct{}

func NewGetArticlesByTagIDUseCase() GetArticlesByTagIDUseCase {
	return &getArticlesByTagIDUseCase{}
}

func (u *getArticlesByTagIDUseCase) Execute(ctx context.Context, tagID string, params gen.GetArticlesByTagIDParams) (*gen.ArticleList, error) {
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
