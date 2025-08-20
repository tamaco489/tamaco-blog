// Package article provides controllers for article-related HTTP endpoints.
// These controllers handle HTTP requests for article operations such as
// retrieving, creating, updating, and deleting articles.
package article

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/gen"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/usecase/article"
)

type UpdateArticleByIDController struct {
	usecase article.UpdateArticleByIDUseCase
}

func NewUpdateArticleByIDController() *UpdateArticleByIDController {
	return &UpdateArticleByIDController{
		usecase: article.NewUpdateArticleByIDUseCase(),
	}
}

func (ctrl *UpdateArticleByIDController) Handle(c *gin.Context, articleID string, req gen.UpdateArticleByIDJSONRequestBody) {
	result, err := ctrl.usecase.UpdateArticleByID(c.Request.Context(), articleID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
