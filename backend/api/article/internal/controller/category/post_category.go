package category

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/gen"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/usecase/category"
)

type CreateCategoryController struct {
	useCase category.CreateCategoryUseCase
}

func NewCreateCategoryController() *CreateCategoryController {
	return &CreateCategoryController{
		useCase: category.NewCreateCategoryUseCase(),
	}
}

func (ctrl *CreateCategoryController) CreateCategory(c *gin.Context) {
	var req gen.CategoryCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := ctrl.useCase.CreateCategory(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, result)
}
