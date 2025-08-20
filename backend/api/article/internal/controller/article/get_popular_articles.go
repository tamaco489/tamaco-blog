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

type GetPopularArticlesController struct {
	usecase article.GetPopularArticlesUseCase
}

func NewGetPopularArticlesController() *GetPopularArticlesController {
	return &GetPopularArticlesController{
		usecase: article.NewGetPopularArticlesUseCase(),
	}
}

func (ctrl *GetPopularArticlesController) Handle(c *gin.Context, params gen.GetPopularArticlesParams) {
	result, err := ctrl.usecase.GetPopularArticles(c.Request.Context(), params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
