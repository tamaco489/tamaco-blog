package article

import (
	"context"
)

// DeleteArticleByIDUseCase handles deleting articles by ID
type DeleteArticleByIDUseCase interface {
	DeleteArticleByID(ctx context.Context, id string) error
}

type deleteArticleByIDUseCase struct {
	// repository層は未実装のため、今回は空
}

// NewDeleteArticleByIDUseCase creates a new delete article by ID usecase
func NewDeleteArticleByIDUseCase() DeleteArticleByIDUseCase {
	return &deleteArticleByIDUseCase{}
}

// DeleteArticleByID implements DeleteArticleByIDUseCase
func (u *deleteArticleByIDUseCase) DeleteArticleByID(ctx context.Context, id string) error {
	// TODO: 実装予定
	return nil
}
