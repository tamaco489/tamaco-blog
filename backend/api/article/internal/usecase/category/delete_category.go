package category

import (
	"context"

	"github.com/google/uuid"
)

// DeleteCategoryUseCase handles deleting a category
type DeleteCategoryUseCase interface {
	DeleteCategory(ctx context.Context, id uuid.UUID) error
}

type deleteCategoryUseCase struct {
	// repository層は未実装のため、今回は空
}

// NewDeleteCategoryUseCase creates a new delete category usecase
func NewDeleteCategoryUseCase() DeleteCategoryUseCase {
	return &deleteCategoryUseCase{}
}

// DeleteCategory implements DeleteCategoryUseCase
func (u *deleteCategoryUseCase) DeleteCategory(ctx context.Context, id uuid.UUID) error {
	// TODO: 実装予定
	return nil
}
