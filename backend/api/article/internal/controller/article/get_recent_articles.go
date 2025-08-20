package article

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/gen"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/usecase/article"
)

// GetRecentArticlesController handles getting recent articles
type GetRecentArticlesController struct {
	usecase article.GetRecentArticlesUseCase
}

// NewGetRecentArticlesController creates a new get recent articles controller
func NewGetRecentArticlesController() *GetRecentArticlesController {
	return &GetRecentArticlesController{
		usecase: article.NewGetRecentArticlesUseCase(),
	}
}

// Handle processes the get recent articles request
func (ctrl *GetRecentArticlesController) Handle(c *gin.Context, params gen.GetRecentArticlesParams) {
	result, err := ctrl.usecase.GetRecentArticles(c.Request.Context(), params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}