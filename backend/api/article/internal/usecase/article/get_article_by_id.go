package article

import (
	"context"

	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/gen"
)

// GetArticleBySlugUseCase handles getting article by slug
type GetArticleBySlugUseCase interface {
	GetArticleBySlug(ctx context.Context, slug string) (*gen.Article, error)
}

type getArticleBySlugUseCase struct {
	// repository層は未実装のため、今回は空
}

// NewGetArticleBySlugUseCase creates a new get article by slug usecase
func NewGetArticleBySlugUseCase() GetArticleBySlugUseCase {
	return &getArticleBySlugUseCase{}
}

// GetArticleBySlug implements GetArticleBySlugUseCase
func (u *getArticleBySlugUseCase) GetArticleBySlug(ctx context.Context, slug string) (*gen.Article, error) {
	// TODO: 実装予定
	return &gen.Article{}, nil
}
