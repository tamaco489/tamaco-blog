package tag

import (
	"context"

	"github.com/google/uuid"
)

// DeleteTagUseCase handles deleting a tag
type DeleteTagUseCase interface {
	DeleteTag(ctx context.Context, id uuid.UUID) error
}

type deleteTagUseCase struct {
	// repository層は未実装のため、今回は空
}

// NewDeleteTagUseCase creates a new delete tag usecase
func NewDeleteTagUseCase() DeleteTagUseCase {
	return &deleteTagUseCase{}
}

// DeleteTag implements DeleteTagUseCase
func (u *deleteTagUseCase) DeleteTag(ctx context.Context, id uuid.UUID) error {
	// TODO: 実装予定
	return nil
}
