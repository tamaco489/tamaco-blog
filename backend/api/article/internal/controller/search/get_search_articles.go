// Package search provides controllers for search-related HTTP endpoints.
// These controllers handle HTTP requests for searching articles
// with various filtering and sorting options.
package search

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/gen"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/usecase/search"
)

type SearchArticlesController struct {
	useCase search.SearchArticlesUseCase
}

func NewSearchArticlesController() *SearchArticlesController {
	return &SearchArticlesController{
		useCase: search.NewSearchArticlesUseCase(),
	}
}

func (ctrl *SearchArticlesController) SearchArticles(c *gin.Context, params gen.SearchArticlesParams) {
	result, err := ctrl.useCase.SearchArticles(c.Request.Context(), params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}
