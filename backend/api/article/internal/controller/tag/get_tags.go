// Package tag provides HTTP controllers for tag endpoints.
package tag

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/usecase/tag"
)

type GetTagsController struct {
	usecase tag.GetTagsUseCase
}

func NewGetTagsController() *GetTagsController {
	return &GetTagsController{
		usecase: tag.NewGetTagsUseCase(),
	}
}

func (ctrl *GetTagsController) GetTags(c *gin.Context) {
	result, err := ctrl.usecase.GetTags(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}
