package tag

import (
	"context"

	"github.com/google/uuid"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/gen"
)

type GetTagByIDUseCase interface {
	GetTagByID(ctx context.Context, tagID string) (*gen.Tag, error)
}

type getTagByIDUseCase struct {
	// repository層は未実装のため、今回は空
}

func NewGetTagByIDUseCase() GetTagByIDUseCase {
	return &getTagByIDUseCase{}
}

func (u *getTagByIDUseCase) GetTagByID(ctx context.Context, tagID string) (*gen.Tag, error) {
	// TODO: 実装予定
	id, _ := uuid.Parse(tagID)
	return &gen.Tag{
		Id:   id,
		Name: "Sample Tag",
	}, nil
}
