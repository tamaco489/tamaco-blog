package article

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/gen"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/usecase/article"
)

// CreateArticleController handles creating new articles
type CreateArticleController struct {
	useCase article.CreateArticleUseCase
}

// NewCreateArticleController creates a new create article controller
func NewCreateArticleController() *CreateArticleController {
	return &CreateArticleController{
		useCase: article.NewCreateArticleUseCase(),
	}
}

// CreateArticle implements gen.ServerInterface
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
