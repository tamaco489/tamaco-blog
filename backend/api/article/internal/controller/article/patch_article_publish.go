package article

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/usecase/article"
)

// PublishArticleByIDController handles publishing an article by ID
type PublishArticleByIDController struct {
	usecase article.PublishArticleByIDUseCase
}

// NewPublishArticleByIDController creates a new publish article by ID controller
func NewPublishArticleByIDController() *PublishArticleByIDController {
	return &PublishArticleByIDController{
		usecase: article.NewPublishArticleByIDUseCase(),
	}
}

// Handle processes the publish article by ID request
func (ctrl *PublishArticleByIDController) Handle(c *gin.Context, articleID string) {
	err := ctrl.usecase.PublishArticleByID(c.Request.Context(), articleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}