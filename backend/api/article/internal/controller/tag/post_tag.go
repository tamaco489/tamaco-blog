// Package tag provides controllers for tag-related HTTP endpoints.
// These controllers handle HTTP requests for tag operations such as
// retrieving, creating, updating, and deleting tags.
package tag

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/gen"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/usecase/tag"
)

type CreateTagController struct {
	useCase tag.CreateTagUseCase
}

func NewCreateTagController() *CreateTagController {
	return &CreateTagController{
		useCase: tag.NewCreateTagUseCase(),
	}
}

func (ctrl *CreateTagController) CreateTag(c *gin.Context) {
	var req gen.TagCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := ctrl.useCase.CreateTag(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, result)
}
