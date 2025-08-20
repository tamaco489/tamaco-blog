package article

import (
	"context"

	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/gen"
)

// GetArticleByIDUseCase handles getting article by ID
type GetArticleByIDUseCase interface {
	GetArticleByID(ctx context.Context, articleID string) (*gen.Article, error)
}

type getArticleByIDUseCase struct {
	// repository層は未実装のため、今回は空
}

// NewGetArticleByIDUseCase creates a new get article by ID usecase
func NewGetArticleByIDUseCase() GetArticleByIDUseCase {
	return &getArticleByIDUseCase{}
}

// GetArticleByID implements GetArticleByIDUseCase
func (u *getArticleByIDUseCase) GetArticleByID(ctx context.Context, articleID string) (*gen.Article, error) {
	// TODO: 実装予定
	return &gen.Article{}, nil
}
