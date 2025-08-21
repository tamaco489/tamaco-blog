package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
	repositoryDS "github.com/tamaco489/tamaco-blog/backend/api/article/internal/repository/data_store"
)

var (
	ErrDatabaseConnection = errors.New("database connection error")
)

// PostgreSQLDB wraps database connection
type PostgreSQLDB struct {
	pool  *pgxpool.Pool
	sqlDB *sql.DB
}

// NewPostgreSQLDB creates a new database wrapper
func NewPostgreSQLDB(ctx context.Context) *PostgreSQLDB {
	// Initialize database connection
	repositoryDS.InitDB(ctx)

	return &PostgreSQLDB{
		pool:  repositoryDS.GetPool(),
		sqlDB: repositoryDS.GetSQLDB(),
	}
}

// GetPool returns the connection pool
func (db *PostgreSQLDB) GetPool() *pgxpool.Pool {
	return db.pool
}

// GetSQLDB returns the database/sql compatible connection
func (db *PostgreSQLDB) GetSQLDB() *sql.DB {
	return db.sqlDB
}
