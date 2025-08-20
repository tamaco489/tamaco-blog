package category

import (
	"context"

	"github.com/google/uuid"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/gen"
)

type GetCategoryByIDUseCase interface {
	GetCategoryByID(ctx context.Context, categoryID string) (*gen.Category, error)
}

type getCategoryByIDUseCase struct {
	// repository層は未実装のため、今回は空
}

func NewGetCategoryByIDUseCase() GetCategoryByIDUseCase {
	return &getCategoryByIDUseCase{}
}

func (u *getCategoryByIDUseCase) GetCategoryByID(ctx context.Context, categoryID string) (*gen.Category, error) {
	// TODO: 実装予定
	id, _ := uuid.Parse(categoryID)
	return &gen.Category{
		Id:   id,
		Name: "Sample Category",
	}, nil
}
