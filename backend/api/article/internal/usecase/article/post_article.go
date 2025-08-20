package article

import (
	"context"

	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/gen"
)

// CreateArticleUseCase handles creating new articles
type CreateArticleUseCase interface {
	CreateArticle(ctx context.Context, req gen.CreateArticleJSONRequestBody) (*gen.Article, error)
}

type createArticleUseCase struct {
	// repository層は未実装のため、今回は空
}

// NewCreateArticleUseCase creates a new create article usecase
func NewCreateArticleUseCase() CreateArticleUseCase {
	return &createArticleUseCase{}
}

// CreateArticle implements CreateArticleUseCase
func (u *createArticleUseCase) CreateArticle(ctx context.Context, req gen.CreateArticleJSONRequestBody) (*gen.Article, error) {
	// TODO: 実装予定
	return &gen.Article{}, nil
}
