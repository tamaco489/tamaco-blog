// Package tag provides tag-related business logic.
package tag

import (
	"context"
	"log/slog"
	"time"

	"github.com/oapi-codegen/runtime/types"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/gen"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/library/config"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/repository"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/repository/gen_sqlc"
)

// GetTagsUseCase handles getting tags list
type GetTagsUseCase interface {
	GetTags(ctx context.Context) (*gen.TagList, error)
}

type getTagsUseCase struct {
	config    *config.Config
	tagRepo   repository.TagRepository
}

// NewGetTagsUseCase creates a new get tags usecase
func NewGetTagsUseCase(cfg *config.Config, tagRepo repository.TagRepository) GetTagsUseCase {
	return &getTagsUseCase{
		config:  cfg,
		tagRepo: tagRepo,
	}
}

// GetTags implements GetTagsUseCase
func (u *getTagsUseCase) GetTags(ctx context.Context) (*gen.TagList, error) {
	slog.InfoContext(ctx, "GetTags called")

	// Retrieve tags from database
	tags, err := u.tagRepo.ListTags(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "failed to list tags", slog.String("error", err.Error()))
		return nil, err
	}

	// Convert database models to API response models
	response := &gen.TagList{
		Tags: make([]gen.Tag, len(tags)),
	}

	for i, tag := range tags {
		response.Tags[i] = convertToGenTag(tag)
	}

	slog.InfoContext(ctx, "successfully retrieved tags", slog.Int("count", len(tags)))
	return response, nil
}

// convertToGenTag converts database model to API response model
func convertToGenTag(dbTag gen_sqlc.Tag) gen.Tag {
	var usageCount *int
	if dbTag.UsageCount.Valid {
		count := int(dbTag.UsageCount.Int32)
		usageCount = &count
	}

	var createdAt *time.Time
	if dbTag.CreatedAt.Valid {
		createdAt = &dbTag.CreatedAt.Time
	}

	var updatedAt *time.Time
	if dbTag.UpdatedAt.Valid {
		updatedAt = &dbTag.UpdatedAt.Time
	}

	return gen.Tag{
		Id:         types.UUID(dbTag.ID),
		Name:       dbTag.Name,
		Slug:       dbTag.Slug,
		UsageCount: usageCount,
		CreatedAt:  createdAt,
		UpdatedAt:  updatedAt,
	}
}
