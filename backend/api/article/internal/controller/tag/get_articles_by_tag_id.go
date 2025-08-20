package tag

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/gen"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/usecase/tag"
)

type GetArticlesByTagIDController struct {
	useCase tag.GetArticlesByTagIDUseCase
}

func NewGetArticlesByTagIDController() *GetArticlesByTagIDController {
	return &GetArticlesByTagIDController{
		useCase: tag.NewGetArticlesByTagIDUseCase(),
	}
}

func (ctrl *GetArticlesByTagIDController) Handle(c *gin.Context, tagID string, params gen.GetArticlesByTagIDParams) {
	// Execute use case
	result, err := ctrl.useCase.Execute(c.Request.Context(), tagID, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
