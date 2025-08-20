package tag

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/gen"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/usecase/tag"
)

type GetArticlesByTagIDController struct {
	usecase tag.GetArticlesByTagIdUseCase
}

func NewGetArticlesByTagIDController() *GetArticlesByTagIDController {
	return &GetArticlesByTagIDController{
		usecase: tag.NewGetArticlesByTagIdUseCase(),
	}
}

func (ctrl *GetArticlesByTagIDController) Handle(c *gin.Context, tagID string, params gen.GetArticlesByTagIdParams) {
	// Execute use case
	result, err := ctrl.usecase.Execute(c.Request.Context(), tagID, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
