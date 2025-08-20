package article

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/usecase/article"
)

// DeleteArticleController handles deleting an article
type DeleteArticleController struct {
	usecase article.DeleteArticleByIDUseCase
}

// NewDeleteArticleController creates a new delete article controller
func NewDeleteArticleController() *DeleteArticleController {
	return &DeleteArticleController{
		usecase: article.NewDeleteArticleByIDUseCase(),
	}
}

// Handle processes the delete article by ID request
func (ctrl *DeleteArticleController) Handle(c *gin.Context, articleID string) {
	err := ctrl.usecase.DeleteArticleByID(c.Request.Context(), articleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}