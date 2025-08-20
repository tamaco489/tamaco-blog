package tag

import (
	"context"

	"github.com/google/uuid"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/gen"
)

// UpdateTagUseCase handles updating a tag
type UpdateTagUseCase interface {
	UpdateTag(ctx context.Context, id uuid.UUID, req gen.TagUpdate) (*gen.Tag, error)
}

type updateTagUseCase struct {
	// repository層は未実装のため、今回は空
}

// NewUpdateTagUseCase creates a new update tag usecase
func NewUpdateTagUseCase() UpdateTagUseCase {
	return &updateTagUseCase{}
}

// UpdateTag implements UpdateTagUseCase
func (u *updateTagUseCase) UpdateTag(ctx context.Context, id uuid.UUID, req gen.TagUpdate) (*gen.Tag, error) {
	// TODO: 実装予定
	return &gen.Tag{
		Id:   id,
		Name: *req.Name,
		Slug: *req.Slug,
	}, nil
}
