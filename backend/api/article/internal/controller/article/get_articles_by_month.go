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

type GetArticlesByMonthController struct {
	usecase article.GetArticlesByMonthUseCase
}

func NewGetArticlesByMonthController() *GetArticlesByMonthController {
	return &GetArticlesByMonthController{
		usecase: article.NewGetArticlesByMonthUseCase(),
	}
}

func (ctrl *GetArticlesByMonthController) Handle(c *gin.Context, year int, month int, params gen.GetArticlesByMonthParams) {
	result, err := ctrl.usecase.GetArticlesByMonth(c.Request.Context(), year, month, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
