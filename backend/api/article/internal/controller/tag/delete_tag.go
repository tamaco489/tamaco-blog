// Package tag provides controllers for tag-related HTTP endpoints.
// These controllers handle HTTP requests for tag operations such as
// retrieving, creating, updating, and deleting tags.
package tag

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/usecase/tag"
)

type DeleteTagController struct {
	useCase tag.DeleteTagUseCase
}

func NewDeleteTagController() *DeleteTagController {
	return &DeleteTagController{
		useCase: tag.NewDeleteTagUseCase(),
	}
}

func (ctrl *DeleteTagController) Handle(c *gin.Context, tagID string) {
	id, err := uuid.Parse(tagID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid tag ID"})
		return
	}

	err = ctrl.useCase.DeleteTag(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
