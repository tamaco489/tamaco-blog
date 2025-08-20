package tag

import (
	"context"

	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/gen"
)

// CreateTagUseCase handles creating a new tag
type CreateTagUseCase interface {
	CreateTag(ctx context.Context, req gen.TagCreate) (*gen.Tag, error)
}

type createTagUseCase struct {
	// repository層は未実装のため、今回は空
}

// NewCreateTagUseCase creates a new create tag usecase
func NewCreateTagUseCase() CreateTagUseCase {
	return &createTagUseCase{}
}

// CreateTag implements CreateTagUseCase
func (u *createTagUseCase) CreateTag(ctx context.Context, req gen.TagCreate) (*gen.Tag, error) {
	// TODO: 実装予定
	return &gen.Tag{
		Name: req.Name,
		Slug: req.Slug,
	}, nil
}
