package category

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/gen"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/usecase/category"
)

type GetArticlesByCategoryIDController struct {
	usecase category.GetArticlesByCategoryIdUseCase
}

func NewGetArticlesByCategoryIDController() *GetArticlesByCategoryIDController {
	return &GetArticlesByCategoryIDController{
		usecase: category.NewGetArticlesByCategoryIdUseCase(),
	}
}

func (ctrl *GetArticlesByCategoryIDController) Handle(c *gin.Context, categoryID string, params gen.GetArticlesByCategoryIdParams) {
	// Execute use case
	result, err := ctrl.usecase.Execute(c.Request.Context(), categoryID, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
