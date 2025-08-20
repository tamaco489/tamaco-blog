package article

import (
	"context"
)

// IncrementArticleViewCountUseCase handles incrementing article view count
type IncrementArticleViewCountUseCase interface {
	IncrementArticleViewCount(ctx context.Context, id string) error
}

type incrementArticleViewCountUseCase struct {
	// repository層は未実装のため、今回は空
}

// NewIncrementArticleViewCountUseCase creates a new increment article view count usecase
func NewIncrementArticleViewCountUseCase() IncrementArticleViewCountUseCase {
	return &incrementArticleViewCountUseCase{}
}

// IncrementArticleViewCount implements IncrementArticleViewCountUseCase
func (u *incrementArticleViewCountUseCase) IncrementArticleViewCount(ctx context.Context, id string) error {
	// TODO: 実装予定
	return nil
}
