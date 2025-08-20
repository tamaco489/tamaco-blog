// Package category provides controllers for category-related HTTP endpoints.
// These controllers handle HTTP requests for category operations such as
// retrieving, creating, updating, and deleting categories.
package category

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/gen"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/usecase/category"
)

type UpdateCategoryController struct {
	useCase category.UpdateCategoryUseCase
}

func NewUpdateCategoryController() *UpdateCategoryController {
	return &UpdateCategoryController{
		useCase: category.NewUpdateCategoryUseCase(),
	}
}

func (ctrl *UpdateCategoryController) Handle(c *gin.Context, categoryID string, req gen.CategoryUpdate) {
	id, err := uuid.Parse(categoryID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid category ID"})
		return
	}

	result, err := ctrl.useCase.UpdateCategory(c.Request.Context(), id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
