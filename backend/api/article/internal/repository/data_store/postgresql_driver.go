package datastore

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
	password := url.QueryEscape(dbConfig.Pass)
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		dbConfig.User, password, dbConfig.Host, dbConfig.Port, dbConfig.Name)
}

var (
	sqlDB *sql.DB
	once  sync.Once
)

func InitDB(ctx context.Context) *sql.DB {
	once.Do(func() {
		appConfig, err := config.GetInstance(ctx)
		if err != nil {
			log.Fatalf("failed to get config instance: %v", err)
		}

		cfg, err := pgxpool.ParseConfig(getDSN(appConfig.CoreDB))
		if err != nil {
			log.Fatalf("failed to parse DB config: %v", err)
		}

		cfg.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol
		cfg.MaxConns = 10
		cfg.MaxConnLifetime = 5 * time.Minute

		sqlDB = stdlib.OpenDB(*cfg.ConnConfig)
		sqlDB.SetMaxOpenConns(int(cfg.MaxConns))
		sqlDB.SetConnMaxLifetime(cfg.MaxConnLifetime)

		if err := sqlDB.PingContext(ctx); err != nil {
			log.Fatalf("failed to ping database: %v", err)
		}
	})
	return sqlDB
}
