// Package category provides controllers for category-related HTTP endpoints.
// These controllers handle HTTP requests for category operations such as
// retrieving, creating, updating, and deleting categories.
package category

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/usecase/category"
)

type GetCategoriesController struct {
	useCase category.GetCategoriesUseCase
}

func NewGetCategoriesController() *GetCategoriesController {
	return &GetCategoriesController{
		useCase: category.NewGetCategoriesUseCase(),
	}
}

func (ctrl *GetCategoriesController) GetCategories(c *gin.Context) {
	result, err := ctrl.useCase.GetCategories(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}
