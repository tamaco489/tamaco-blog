// Package health provides HTTP controllers for health check endpoints.
package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/usecase/health"
)

// HealthcheckController handles health check requests
type HealthcheckController struct {
	usecase health.HealthcheckUseCase
}

// NewHealthcheckController creates a new healthcheck controller
func NewHealthcheckController() *HealthcheckController {
	return &HealthcheckController{
		usecase: health.NewHealthcheckUseCase(),
	}
}

// Healthcheck implements gen.ServerInterface
func (ctrl *HealthcheckController) Healthcheck(c *gin.Context) {
	result, err := ctrl.usecase.Healthcheck(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}