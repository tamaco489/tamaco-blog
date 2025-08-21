// Package handler provides HTTP handlers for the API.
package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/controller"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/gen"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/library/config"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/library/logger"

	datastore "github.com/tamaco489/tamaco-blog/backend/api/article/internal/repository/data_store"
)

func NewHandler(ctx context.Context) (*http.Server, error) {
	cfg, err := config.GetInstance(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get config instance: %w", err)
	}

	db := datastore.InitDB(ctx)

	r := gin.New()
	r.Use(gin.LoggerWithFormatter(logger.GinLogFormatter))
	r.Use(gin.Recovery())

	// CORS設定
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{
		"http://localhost:3000",   // フロントエンド開発環境
		"https://tamaco-blog.com", // 本番環境（予定）
	}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	corsConfig.AllowCredentials = true
	r.Use(cors.New(corsConfig))

	ctrl := controller.NewServerController(ctx, cfg, db)

	gen.RegisterHandlers(r, ctrl)

	srv := &http.Server{
		Addr:    cfg.GetAPIAddr(),
		Handler: r,
	}

	return srv, nil
}
