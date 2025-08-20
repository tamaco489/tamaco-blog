package category

import (
	"context"

	"github.com/google/uuid"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/gen"
)

// UpdateCategoryUseCase handles updating a category
type UpdateCategoryUseCase interface {
	UpdateCategory(ctx context.Context, id uuid.UUID, req gen.CategoryUpdate) (*gen.Category, error)
}

type updateCategoryUseCase struct {
	// repository層は未実装のため、今回は空
}

// NewUpdateCategoryUseCase creates a new update category usecase
func NewUpdateCategoryUseCase() UpdateCategoryUseCase {
	return &updateCategoryUseCase{}
}

// UpdateCategory implements UpdateCategoryUseCase
func (u *updateCategoryUseCase) UpdateCategory(ctx context.Context, id uuid.UUID, req gen.CategoryUpdate) (*gen.Category, error) {
	// TODO: 実装予定
	return &gen.Category{
		Id:   id,
		Name: *req.Name,
		Slug: *req.Slug,
	}, nil
}
