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

type CreateArticleController struct {
	useCase article.CreateArticleUseCase
}

func NewCreateArticleController() *CreateArticleController {
	return &CreateArticleController{
		useCase: article.NewCreateArticleUseCase(),
	}
}

func (ctrl *CreateArticleController) CreateArticle(c *gin.Context) {
	var req gen.CreateArticleJSONRequestBody
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := ctrl.useCase.CreateArticle(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, result)
}
