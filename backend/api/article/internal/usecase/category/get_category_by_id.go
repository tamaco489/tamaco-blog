package category

import (
	"context"

	"github.com/google/uuid"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/gen"
)

// GetCategoryByIdUseCase handles getting a category by ID
type GetCategoryByIdUseCase interface {
	GetCategoryById(ctx context.Context, categoryId string) (*gen.Category, error)
}

type getCategoryByIdUseCase struct {
	// repository層は未実装のため、今回は空
}

// NewGetCategoryByIdUseCase creates a new get category by ID usecase
func NewGetCategoryByIdUseCase() GetCategoryByIdUseCase {
	return &getCategoryByIdUseCase{}
}

// GetCategoryById implements GetCategoryByIdUseCase
func (u *getCategoryByIdUseCase) GetCategoryById(ctx context.Context, categoryId string) (*gen.Category, error) {
	// TODO: 実装予定
	id, _ := uuid.Parse(categoryId)
	return &gen.Category{
		Id: id,
		Name: "Sample Category",
	}, nil
}
