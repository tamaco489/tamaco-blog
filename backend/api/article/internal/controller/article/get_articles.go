// Package article provides controllers for article-related HTTP endpoints.
// These controllers handle HTTP requests for article operations such as
// retrieving, creating, updating, and deleting articles.
package article

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/gen"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/usecase/article"
)

type GetArticlesController struct {
	useCase article.GetArticlesUseCase
}

func NewGetArticlesController() *GetArticlesController {
	return &GetArticlesController{
		useCase: article.NewGetArticlesUseCase(),
	}
}

func (ctrl *GetArticlesController) GetArticles(c *gin.Context, params gen.GetArticlesParams) {
	result, err := ctrl.useCase.GetArticles(c.Request.Context(), params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}
