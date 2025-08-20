package article

import (
	"context"
)

// PublishArticleByIDUseCase handles publishing articles by ID
type PublishArticleByIDUseCase interface {
	PublishArticleByID(ctx context.Context, id string) error
}

type publishArticleByIDUseCase struct {
	// repository層は未実装のため、今回は空
}

// NewPublishArticleByIDUseCase creates a new publish article by ID usecase
func NewPublishArticleByIDUseCase() PublishArticleByIDUseCase {
	return &publishArticleByIDUseCase{}
}

// PublishArticleByID implements PublishArticleByIDUseCase
func (u *publishArticleByIDUseCase) PublishArticleByID(ctx context.Context, id string) error {
	// TODO: 実装予定
	return nil
}
