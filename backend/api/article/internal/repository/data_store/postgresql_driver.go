package repository

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
	SSLMode  string
}

func GetDBConfig() DBConfig {
	return DBConfig{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSL_MODE"),
	}
}

func getDSN(cfg DBConfig) string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode)
}

var (
	dbPool *pgxpool.Pool
	once   sync.Once
)

func InitDB(ctx context.Context) {
	once.Do(func() {
		cfg, err := pgxpool.ParseConfig(getDSN(GetDBConfig()))
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
	})
}

func GetPool() *pgxpool.Pool {
	return dbPool
}
