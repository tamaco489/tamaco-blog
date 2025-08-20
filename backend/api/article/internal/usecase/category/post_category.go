package category

import (
	"context"

	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/gen"
)

// CreateCategoryUseCase handles creating a new category
type CreateCategoryUseCase interface {
	CreateCategory(ctx context.Context, req gen.CategoryCreate) (*gen.Category, error)
}

type createCategoryUseCase struct {
	// repository層は未実装のため、今回は空
}

// NewCreateCategoryUseCase creates a new create category usecase
func NewCreateCategoryUseCase() CreateCategoryUseCase {
	return &createCategoryUseCase{}
}

// CreateCategory implements CreateCategoryUseCase
func (u *createCategoryUseCase) CreateCategory(ctx context.Context, req gen.CategoryCreate) (*gen.Category, error) {
	// TODO: 実装予定
	return &gen.Category{
		Name: req.Name,
		Slug: req.Slug,
	}, nil
}
