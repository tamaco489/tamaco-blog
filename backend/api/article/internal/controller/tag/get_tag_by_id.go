// Package tag provides controllers for tag-related HTTP endpoints.
// These controllers handle HTTP requests for tag operations such as
// retrieving, creating, updating, and deleting tags.
package tag

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/usecase/tag"
)

type GetTagByIDController struct {
	useCase tag.GetTagByIDUseCase
}

func NewGetTagByIDController() *GetTagByIDController {
	return &GetTagByIDController{
		useCase: tag.NewGetTagByIDUseCase(),
	}
}

func (ctrl *GetTagByIDController) Handle(c *gin.Context, tagID string) {
	result, err := ctrl.useCase.GetTagByID(c.Request.Context(), tagID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}
