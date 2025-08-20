package article

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/usecase/article"
)

// GetArticleByIDController handles getting an article by ID
type GetArticleByIDController struct {
	usecase article.GetArticleByIDUseCase
}

// NewGetArticleByIDController creates a new get article by ID controller
func NewGetArticleByIDController() *GetArticleByIDController {
	return &GetArticleByIDController{
		usecase: article.NewGetArticleByIDUseCase(),
	}
}

// Handle processes the get article by ID request
func (ctrl *GetArticleByIDController) Handle(c *gin.Context, articleID string) {
	result, err := ctrl.usecase.GetArticleByID(c.Request.Context(), articleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}