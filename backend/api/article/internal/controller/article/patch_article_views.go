// Package article provides controllers for article-related HTTP endpoints.
// These controllers handle HTTP requests for article operations such as
// retrieving, creating, updating, and deleting articles.
package article

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/usecase/article"
)

type IncrementArticleViewCountController struct {
	usecase article.IncrementArticleViewCountUseCase
}

func NewIncrementArticleViewCountController() *IncrementArticleViewCountController {
	return &IncrementArticleViewCountController{
		usecase: article.NewIncrementArticleViewCountUseCase(),
	}
}

func (ctrl *IncrementArticleViewCountController) Handle(c *gin.Context, articleID string) {
	err := ctrl.usecase.IncrementArticleViewCount(c.Request.Context(), articleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
