package category

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/usecase/category"
)

type DeleteCategoryController struct {
	useCase category.DeleteCategoryUseCase
}

func NewDeleteCategoryController() *DeleteCategoryController {
	return &DeleteCategoryController{
		useCase: category.NewDeleteCategoryUseCase(),
	}
}

func (ctrl *DeleteCategoryController) Handle(c *gin.Context, categoryID string) {
	id, err := uuid.Parse(categoryID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid category ID"})
		return
	}

	err = ctrl.useCase.DeleteCategory(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
