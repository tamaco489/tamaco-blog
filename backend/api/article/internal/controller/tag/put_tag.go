package tag

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/gen"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/usecase/tag"
)

type UpdateTagController struct {
	useCase tag.UpdateTagUseCase
}

func NewUpdateTagController() *UpdateTagController {
	return &UpdateTagController{
		useCase: tag.NewUpdateTagUseCase(),
	}
}

func (ctrl *UpdateTagController) Handle(c *gin.Context, tagID string, req gen.TagUpdate) {
	id, err := uuid.Parse(tagID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid tag ID"})
		return
	}

	result, err := ctrl.useCase.UpdateTag(c.Request.Context(), id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
