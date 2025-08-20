package category

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/gen"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/usecase/category"
)

type GetArticlesByCategoryIDController struct {
	useCase category.GetArticlesByCategoryIDUseCase
}

func NewGetArticlesByCategoryIDController() *GetArticlesByCategoryIDController {
	return &GetArticlesByCategoryIDController{
		useCase: category.NewGetArticlesByCategoryIDUseCase(),
	}
}

func (ctrl *GetArticlesByCategoryIDController) Handle(c *gin.Context, categoryID string, params gen.GetArticlesByCategoryIDParams) {
	// Execute use case
	result, err := ctrl.useCase.Execute(c.Request.Context(), categoryID, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
