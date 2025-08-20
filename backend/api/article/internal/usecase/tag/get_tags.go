// Package tag provides tag-related business logic.
package tag

import (
	"context"
	"log/slog"

	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/gen"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/library/config"
)

// GetTagsUseCase handles getting tags list
type GetTagsUseCase interface {
	GetTags(ctx context.Context) (*gen.TagList, error)
}

type getTagsUseCase struct {
	config *config.Config
}

// NewGetTagsUseCase creates a new get tags usecase
func NewGetTagsUseCase(cfg *config.Config) GetTagsUseCase {
	return &getTagsUseCase{
		config: cfg,
	}
}

// GetTags implements GetTagsUseCase
func (u *getTagsUseCase) GetTags(ctx context.Context) (*gen.TagList, error) {

	slog.InfoContext(ctx, "[TEST] GetTags called")

	// DB接続情報をログ出力（デバッグ用）
	slog.InfoContext(ctx, "[DEBUG] Database connection info",
		slog.String("host", u.config.CoreDB.Host),
		slog.String("port", u.config.CoreDB.Port),
		slog.String("user", u.config.CoreDB.User),
		slog.String("dbname", u.config.CoreDB.Name),
		slog.String("dsn", u.config.GetDatabaseDSN()),
		slog.Bool("is_production", u.config.IsProduction()),
		slog.String("api_env", string(u.config.API.Env)),
	)

	// TODO: 実装予定
	return &gen.TagList{
		Tags: []gen.Tag{},
	}, nil
}
