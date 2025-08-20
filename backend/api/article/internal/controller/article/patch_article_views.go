package article

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/usecase/article"
)

// IncrementArticleViewCountController handles incrementing an article view count
type IncrementArticleViewCountController struct {
	usecase article.IncrementArticleViewCountUseCase
}

// NewIncrementArticleViewCountController creates a new increment article view count controller
func NewIncrementArticleViewCountController() *IncrementArticleViewCountController {
	return &IncrementArticleViewCountController{
		usecase: article.NewIncrementArticleViewCountUseCase(),
	}
}

// Handle processes the increment article view count request
func (ctrl *IncrementArticleViewCountController) Handle(c *gin.Context, articleID string) {
	err := ctrl.usecase.IncrementArticleViewCount(c.Request.Context(), articleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}