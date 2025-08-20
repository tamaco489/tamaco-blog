// Package article provides controllers for article-related HTTP endpoints.
// These controllers handle HTTP requests for article operations such as
// retrieving, creating, updating, and deleting articles.
package article

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/usecase/article"
)

type GetArticleByIDController struct {
	usecase article.GetArticleByIDUseCase
}

func NewGetArticleByIDController() *GetArticleByIDController {
	return &GetArticleByIDController{
		usecase: article.NewGetArticleByIDUseCase(),
	}
}

func (ctrl *GetArticleByIDController) Handle(c *gin.Context, articleID string) {
	result, err := ctrl.usecase.GetArticleByID(c.Request.Context(), articleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
