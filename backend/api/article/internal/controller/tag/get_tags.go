// Package tag provides controllers for tag-related HTTP endpoints.
// These controllers handle HTTP requests for tag operations such as
// retrieving, creating, updating, and deleting tags.
package tag

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/library/config"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/usecase/tag"
)

type GetTagsController struct {
	useCase tag.GetTagsUseCase
}

func NewGetTagsController(ctx context.Context, cfg *config.Config) *GetTagsController {
	return &GetTagsController{
		useCase: tag.NewGetTagsUseCase(cfg),
	}
}

func (ctrl *GetTagsController) GetTags(c *gin.Context) {
	result, err := ctrl.useCase.GetTags(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}
