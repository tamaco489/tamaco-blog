package article

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/gen"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/usecase/article"
)

// UpdateArticleByIDController handles updating an article by ID
type UpdateArticleByIDController struct {
	usecase article.UpdateArticleByIDUseCase
}

// NewUpdateArticleByIDController creates a new update article by ID controller
func NewUpdateArticleByIDController() *UpdateArticleByIDController {
	return &UpdateArticleByIDController{
		usecase: article.NewUpdateArticleByIDUseCase(),
	}
}

// Handle processes the update article by ID request
func (ctrl *UpdateArticleByIDController) Handle(c *gin.Context, articleID string, req gen.UpdateArticleByIDJSONRequestBody) {
	result, err := ctrl.usecase.UpdateArticleByID(c.Request.Context(), articleID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}