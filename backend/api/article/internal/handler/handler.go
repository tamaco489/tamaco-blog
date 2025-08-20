// Package handler provides HTTP handlers for the API.
package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/controller"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/gen"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/library/config"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/library/logger"
)

func NewHandler(ctx context.Context) (*http.Server, error) {
	cfg, err := config.GetInstance(ctx)
	if err != nil {
		return nil, err
	}

	r := gin.New()
	r.Use(gin.LoggerWithFormatter(logger.GinLogFormatter))
	r.Use(gin.Recovery())

	ctrl := controller.NewServerController()

	gen.RegisterHandlers(r, ctrl)

	srv := &http.Server{
		Addr:    cfg.GetAPIAddr(),
		Handler: r,
	}

	return srv, nil
}
