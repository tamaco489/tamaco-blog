// Package search provides HTTP controllers for search endpoints.
package search

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/gen"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/usecase/search"
)

// SearchArticlesController handles searching articles
type SearchArticlesController struct {
	useCase search.SearchArticlesUseCase
}

// NewSearchArticlesController creates a new search articles controller
func NewSearchArticlesController() *SearchArticlesController {
	return &SearchArticlesController{
		useCase: search.NewSearchArticlesUseCase(),
	}
}

// SearchArticles implements gen.ServerInterface
func (ctrl *SearchArticlesController) SearchArticles(c *gin.Context, params gen.SearchArticlesParams) {
	result, err := ctrl.useCase.SearchArticles(c.Request.Context(), params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}