package category

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/gen"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/usecase/category"
)

// UpdateCategoryController handles updating a category
type UpdateCategoryController struct {
	usecase category.UpdateCategoryUseCase
}

// NewUpdateCategoryController creates a new update category controller
func NewUpdateCategoryController() *UpdateCategoryController {
	return &UpdateCategoryController{
		usecase: category.NewUpdateCategoryUseCase(),
	}
}

// Handle processes the update category by ID request
func (ctrl *UpdateCategoryController) Handle(c *gin.Context, categoryID string, req gen.CategoryUpdate) {
	id, err := uuid.Parse(categoryID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid category ID"})
		return
	}

	result, err := ctrl.usecase.UpdateCategory(c.Request.Context(), id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
