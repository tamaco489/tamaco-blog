package article

import (
	"context"

	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/gen"
)

// UpdateArticleByIDUseCase handles updating articles by ID
type UpdateArticleByIDUseCase interface {
	UpdateArticleByID(ctx context.Context, id string, req gen.UpdateArticleByIDJSONRequestBody) (*gen.Article, error)
}

type updateArticleByIDUseCase struct {
	// repository層は未実装のため、今回は空
}

// NewUpdateArticleByIDUseCase creates a new update article by ID usecase
func NewUpdateArticleByIDUseCase() UpdateArticleByIDUseCase {
	return &updateArticleByIDUseCase{}
}

// UpdateArticleByID implements UpdateArticleByIDUseCase
func (u *updateArticleByIDUseCase) UpdateArticleByID(ctx context.Context, id string, req gen.UpdateArticleByIDJSONRequestBody) (*gen.Article, error) {
	// TODO: 実装予定
	return &gen.Article{}, nil
}
