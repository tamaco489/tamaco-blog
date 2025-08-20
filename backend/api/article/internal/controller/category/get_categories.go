// Package category provides HTTP controllers for category endpoints.
package category

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/usecase/category"
)

type GetCategoriesController struct {
	usecase category.GetCategoriesUseCase
}

func NewGetCategoriesController() *GetCategoriesController {
	return &GetCategoriesController{
		usecase: category.NewGetCategoriesUseCase(),
	}
}

func (ctrl *GetCategoriesController) GetCategories(c *gin.Context) {
	result, err := ctrl.usecase.GetCategories(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}
