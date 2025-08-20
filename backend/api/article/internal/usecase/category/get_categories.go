// Package category provides category-related business logic.
package category

import (
	"context"
	"log/slog"

	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/gen"
)

// GetCategoriesUseCase handles getting categories list
type GetCategoriesUseCase interface {
	GetCategories(ctx context.Context) (*gen.CategoryList, error)
}

type getCategoriesUseCase struct {
	// repository層は未実装のため、今回は空
}

// NewGetCategoriesUseCase creates a new get categories usecase
func NewGetCategoriesUseCase() GetCategoriesUseCase {
	return &getCategoriesUseCase{}
}

// GetCategories implements GetCategoriesUseCase
func (u *getCategoriesUseCase) GetCategories(ctx context.Context) (*gen.CategoryList, error) {

	slog.InfoContext(ctx, "[TEST] GetCategories called")

	// TODO: 実装予定
	return &gen.CategoryList{
		Categories: []gen.Category{},
	}, nil
}
