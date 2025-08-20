// Package health provides controllers for health check endpoints.
// This package implements health monitoring functionality to verify
// the application's operational status.
package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/usecase/health"
)

type HealthcheckController struct {
	useCase health.HealthcheckUseCase
}

func NewHealthcheckController() *HealthcheckController {
	return &HealthcheckController{
		useCase: health.NewHealthcheckUseCase(),
	}
}

func (ctrl *HealthcheckController) Healthcheck(c *gin.Context) {
	result, err := ctrl.useCase.Healthcheck(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}
