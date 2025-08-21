package repository

import (
	"context"
	"log"

	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/repository/gen_sqlc"
)

// TagRepository provides tag-related database operations
type TagRepository interface {
	ListTags(ctx context.Context) ([]gen_sqlc.Tag, error)
}

type tagRepository struct {
	db *PostgreSQLDB
}

// NewTagRepository creates a new tag repository
func NewTagRepository(db *PostgreSQLDB) TagRepository {
	return &tagRepository{
		db: db,
	}
}

// ListTags retrieves all tags from database
func (r *tagRepository) ListTags(ctx context.Context) ([]gen_sqlc.Tag, error) {
	sqlDB := r.db.GetSQLDB()
	if sqlDB == nil {
		log.Printf("database connection is nil")
		return nil, ErrDatabaseConnection
	}

	queries := gen_sqlc.New()
	tags, err := queries.ListTags(ctx, sqlDB)
	if err != nil {
		log.Printf("failed to list tags: %v", err)
		return nil, err
	}

	return tags, nil
}