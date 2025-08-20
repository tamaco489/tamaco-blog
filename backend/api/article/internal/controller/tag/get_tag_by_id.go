package tag

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/usecase/tag"
)

type GetTagByIDController struct {
	usecase tag.GetTagByIdUseCase
}

func NewGetTagByIDController() *GetTagByIDController {
	return &GetTagByIDController{
		usecase: tag.NewGetTagByIdUseCase(),
	}
}

func (ctrl *GetTagByIDController) Handle(c *gin.Context, tagID string) {
	result, err := ctrl.usecase.GetTagById(c.Request.Context(), tagID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}
