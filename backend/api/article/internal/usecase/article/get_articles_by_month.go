package article

import (
	"context"

	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/gen"
)

// GetArticlesByMonthUseCase handles getting articles by month
type GetArticlesByMonthUseCase interface {
	GetArticlesByMonth(ctx context.Context, year int, month int, params gen.GetArticlesByMonthParams) (*gen.ArticleList, error)
}

type getArticlesByMonthUseCase struct {
	// repository層は未実装のため、今回は空
}

// NewGetArticlesByMonthUseCase creates a new get articles by month usecase
func NewGetArticlesByMonthUseCase() GetArticlesByMonthUseCase {
	return &getArticlesByMonthUseCase{}
}

// GetArticlesByMonth implements GetArticlesByMonthUseCase
func (u *getArticlesByMonthUseCase) GetArticlesByMonth(ctx context.Context, year int, month int, params gen.GetArticlesByMonthParams) (*gen.ArticleList, error) {
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
