package article

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/gen"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/usecase/article"
)

// GetArticlesByMonthController handles getting articles by month
type GetArticlesByMonthController struct {
	usecase article.GetArticlesByMonthUseCase
}

// NewGetArticlesByMonthController creates a new get articles by month controller
func NewGetArticlesByMonthController() *GetArticlesByMonthController {
	return &GetArticlesByMonthController{
		usecase: article.NewGetArticlesByMonthUseCase(),
	}
}

// Handle processes the get articles by month request
func (ctrl *GetArticlesByMonthController) Handle(c *gin.Context, year int, month int, params gen.GetArticlesByMonthParams) {
	result, err := ctrl.usecase.GetArticlesByMonth(c.Request.Context(), year, month, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}