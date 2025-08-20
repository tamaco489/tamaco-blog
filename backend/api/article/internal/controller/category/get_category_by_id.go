package category

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/usecase/category"
)

type GetCategoryByIDController struct {
	useCase category.GetCategoryByIDUseCase
}

func NewGetCategoryByIDController() *GetCategoryByIDController {
	return &GetCategoryByIDController{
		useCase: category.NewGetCategoryByIDUseCase(),
	}
}

func (ctrl *GetCategoryByIDController) Handle(c *gin.Context, categoryID string) {
	result, err := ctrl.useCase.GetCategoryByID(c.Request.Context(), categoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}
