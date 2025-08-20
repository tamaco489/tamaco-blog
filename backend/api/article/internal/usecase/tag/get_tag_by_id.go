package tag

import (
	"context"

	"github.com/google/uuid"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/gen"
)

// GetTagByIdUseCase handles getting a tag by ID
type GetTagByIdUseCase interface {
	GetTagById(ctx context.Context, tagId string) (*gen.Tag, error)
}

type getTagByIdUseCase struct {
	// repository層は未実装のため、今回は空
}

// NewGetTagByIdUseCase creates a new get tag by ID usecase
func NewGetTagByIdUseCase() GetTagByIdUseCase {
	return &getTagByIdUseCase{}
}

// GetTagById implements GetTagByIdUseCase
func (u *getTagByIdUseCase) GetTagById(ctx context.Context, tagId string) (*gen.Tag, error) {
	// TODO: 実装予定
	id, _ := uuid.Parse(tagId)
	return &gen.Tag{
		Id: id,
		Name: "Sample Tag",
	}, nil
}
