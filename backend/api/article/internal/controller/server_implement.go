package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/gen"
)

// GetArticles implements gen.ServerInterface
func (s *ServerController) GetArticles(c *gin.Context, params gen.GetArticlesParams) {
	s.getArticlesCtrl.GetArticles(c, params)
}

// CreateArticle implements gen.ServerInterface
func (s *ServerController) CreateArticle(c *gin.Context) {
	s.createArticleCtrl.CreateArticle(c)
}

// Healthcheck implements gen.ServerInterface
func (s *ServerController) Healthcheck(c *gin.Context) {
	s.healthcheckCtrl.Healthcheck(c)
}

func (s *ServerController) GetArticlesByMonth(c *gin.Context, year gen.YearParam, month gen.MonthParam, params gen.GetArticlesByMonthParams) {
	s.getArticlesByMonthCtrl.Handle(c, int(year), int(month), params)
}

func (s *ServerController) GetPopularArticles(c *gin.Context, params gen.GetPopularArticlesParams) {
	s.getPopularArticlesCtrl.Handle(c, params)
}

func (s *ServerController) GetRecentArticles(c *gin.Context, params gen.GetRecentArticlesParams) {
	s.getRecentArticlesCtrl.Handle(c, params)
}

func (s *ServerController) DeleteArticleByID(c *gin.Context, articleID string) {
	s.deleteArticleCtrl.Handle(c, articleID)
}

func (s *ServerController) UpdateArticleByID(c *gin.Context, articleID string) {
	var req gen.UpdateArticleByIDJSONRequestBody
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	s.updateArticleByIDCtrl.Handle(c, articleID, req)
}

func (s *ServerController) PublishArticleByID(c *gin.Context, articleID string) {
	s.publishArticleByIDCtrl.Handle(c, articleID)
}

func (s *ServerController) IncrementArticleViewCount(c *gin.Context, articleID string) {
	s.incrementViewCountCtrl.Handle(c, articleID)
}

func (s *ServerController) GetArticleByID(c *gin.Context, articleID string) {
	s.getArticleByIDCtrl.Handle(c, articleID)
}

func (s *ServerController) GetCategories(c *gin.Context) {
	s.getCategoriesCtrl.GetCategories(c)
}

func (s *ServerController) CreateCategory(c *gin.Context) {
	s.createCategoryCtrl.CreateCategory(c)
}

func (s *ServerController) DeleteCategoryByID(c *gin.Context, categoryID string) {
	s.deleteCategoryCtrl.Handle(c, categoryID)
}

func (s *ServerController) UpdateCategoryByID(c *gin.Context, categoryID string) {
	var req gen.CategoryUpdate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	s.updateCategoryCtrl.Handle(c, categoryID, req)
}

func (s *ServerController) GetCategoryByID(c *gin.Context, categoryID string) {
	s.getCategoryByIDCtrl.Handle(c, categoryID)
}

func (s *ServerController) GetArticlesByCategoryID(c *gin.Context, categoryID string, params gen.GetArticlesByCategoryIDParams) {
	s.getArticlesByCategoryIDCtrl.Handle(c, categoryID, params)
}

func (s *ServerController) SearchArticles(c *gin.Context, params gen.SearchArticlesParams) {
	s.searchArticlesCtrl.SearchArticles(c, params)
}

func (s *ServerController) GetTags(c *gin.Context) {
	s.getTagsCtrl.GetTags(c)
}

func (s *ServerController) CreateTag(c *gin.Context) {
	s.createTagCtrl.CreateTag(c)
}

func (s *ServerController) DeleteTagByID(c *gin.Context, tagID string) {
	s.deleteTagCtrl.Handle(c, tagID)
}

func (s *ServerController) UpdateTagByID(c *gin.Context, tagID string) {
	var req gen.TagUpdate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	s.updateTagCtrl.Handle(c, tagID, req)
}

func (s *ServerController) GetTagByID(c *gin.Context, tagID string) {
	s.getTagByIDCtrl.Handle(c, tagID)
}

func (s *ServerController) GetArticlesByTagID(c *gin.Context, tagID string, params gen.GetArticlesByTagIDParams) {
	s.getArticlesByTagIDCtrl.Handle(c, tagID, params)
}
