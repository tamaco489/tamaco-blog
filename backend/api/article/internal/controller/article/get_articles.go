// Package article provides HTTP controllers for article endpoints.
package article

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/gen"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/usecase/article"
)

// GetArticlesController handles getting articles list
type GetArticlesController struct {
	usecase article.GetArticlesUseCase
}

// NewGetArticlesController creates a new get articles controller
func NewGetArticlesController() *GetArticlesController {
	return &GetArticlesController{
		usecase: article.NewGetArticlesUseCase(),
	}
}

// GetArticles implements gen.ServerInterface
func (ctrl *GetArticlesController) GetArticles(c *gin.Context, params gen.GetArticlesParams) {
	result, err := ctrl.usecase.GetArticles(c.Request.Context(), params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}
