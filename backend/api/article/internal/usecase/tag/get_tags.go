// Package tag provides tag-related business logic.
package tag

import (
	"context"
	"log/slog"

	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/gen"
)

// GetTagsUseCase handles getting tags list
type GetTagsUseCase interface {
	GetTags(ctx context.Context) (*gen.TagList, error)
}

type getTagsUseCase struct {
	// repository層は未実装のため、今回は空
}

// NewGetTagsUseCase creates a new get tags usecase
func NewGetTagsUseCase() GetTagsUseCase {
	return &getTagsUseCase{}
}

// GetTags implements GetTagsUseCase
func (u *getTagsUseCase) GetTags(ctx context.Context) (*gen.TagList, error) {

	slog.InfoContext(ctx, "[TEST] GetTags called")

	// TODO: 実装予定
	return &gen.TagList{
		Tags: []gen.Tag{},
	}, nil
}
