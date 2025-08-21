package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"sync"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/library/config"
)

func getDSN(dbConfig config.DatabaseConfig) string {
	// URL encode password to handle special characters
	password := url.QueryEscape(dbConfig.Pass)
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		dbConfig.User, password, dbConfig.Host, dbConfig.Port, dbConfig.Name)
}

var (
	dbPool *pgxpool.Pool
	sqlDB  *sql.DB
	once   sync.Once
)

func InitDB(ctx context.Context) {
	once.Do(func() {
		// Get config from singleton
		appConfig, err := config.GetInstance(ctx)
		if err != nil {
			log.Fatalf("failed to get config instance: %v", err)
		}

		// Parse database connection string
		cfg, err := pgxpool.ParseConfig(getDSN(appConfig.CoreDB))
		if err != nil {
			log.Fatalf("failed to parse DB config: %v", err)
		}

		cfg.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol
		cfg.MaxConns = 2
		cfg.MaxConnLifetime = 1 * time.Minute
		dbPool, err = pgxpool.NewWithConfig(ctx, cfg)
		if err != nil {
			log.Fatalf("failed to connect to database: %v", err)
		}

		// Create database/sql compatible connection
		connConfig := cfg.ConnConfig
		sqlDB = stdlib.OpenDB(*connConfig)
		sqlDB.SetMaxOpenConns(int(cfg.MaxConns))
		sqlDB.SetConnMaxLifetime(cfg.MaxConnLifetime)
	})
}

func GetPool() *pgxpool.Pool {
	return dbPool
}

func GetSQLDB() *sql.DB {
	return sqlDB
}
