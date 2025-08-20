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

type GetRecentArticlesController struct {
	usecase article.GetRecentArticlesUseCase
}

func NewGetRecentArticlesController() *GetRecentArticlesController {
	return &GetRecentArticlesController{
		usecase: article.NewGetRecentArticlesUseCase(),
	}
}

func (ctrl *GetRecentArticlesController) Handle(c *gin.Context, params gen.GetRecentArticlesParams) {
	result, err := ctrl.usecase.GetRecentArticles(c.Request.Context(), params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
