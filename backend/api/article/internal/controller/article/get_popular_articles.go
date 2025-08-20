package article

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/gen"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/usecase/article"
)

// GetPopularArticlesController handles getting popular articles
type GetPopularArticlesController struct {
	usecase article.GetPopularArticlesUseCase
}

// NewGetPopularArticlesController creates a new get popular articles controller
func NewGetPopularArticlesController() *GetPopularArticlesController {
	return &GetPopularArticlesController{
		usecase: article.NewGetPopularArticlesUseCase(),
	}
}

// Handle processes the get popular articles request
func (ctrl *GetPopularArticlesController) Handle(c *gin.Context, params gen.GetPopularArticlesParams) {
	result, err := ctrl.usecase.GetPopularArticles(c.Request.Context(), params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}