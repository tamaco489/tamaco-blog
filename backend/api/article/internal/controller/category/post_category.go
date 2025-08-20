package category

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/gen"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/usecase/category"
)

// CreateCategoryController handles creating a new category
type CreateCategoryController struct {
	usecase category.CreateCategoryUseCase
}

// NewCreateCategoryController creates a new create category controller
func NewCreateCategoryController() *CreateCategoryController {
	return &CreateCategoryController{
		usecase: category.NewCreateCategoryUseCase(),
	}
}

// CreateCategory implements gen.ServerInterface
func (ctrl *CreateCategoryController) CreateCategory(c *gin.Context) {
	var req gen.CategoryCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := ctrl.usecase.CreateCategory(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, result)
}
