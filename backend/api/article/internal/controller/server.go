// Package controller provides HTTP controllers for the API.
package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/controller/article"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/controller/category"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/controller/health"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/controller/tag"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/gen"
)

// ServerController aggregates all controllers and implements gen.ServerInterface
type ServerController struct {
	// Article controllers
	getArticlesCtrl             *article.GetArticlesController
	createArticleCtrl           *article.CreateArticleController
	getArticlesByTagIDCtrl      *tag.GetArticlesByTagIDController
	getArticlesByCategoryIDCtrl *category.GetArticlesByCategoryIDController

	// Health controller
	healthcheckCtrl *health.HealthcheckController

	// Category controllers
	getCategoriesCtrl   *category.GetCategoriesController
	createCategoryCtrl  *category.CreateCategoryController
	deleteCategoryCtrl  *category.DeleteCategoryController
	updateCategoryCtrl  *category.UpdateCategoryController
	getCategoryByIDCtrl *category.GetCategoryByIDController

	// Tag controllers
	getTagsCtrl    *tag.GetTagsController
	createTagCtrl  *tag.CreateTagController
	deleteTagCtrl  *tag.DeleteTagController
	updateTagCtrl  *tag.UpdateTagController
	getTagByIDCtrl *tag.GetTagByIDController
}

// NewServerController creates a new server controller
func NewServerController() *ServerController {
	return &ServerController{
		// Health controller
		healthcheckCtrl: health.NewHealthcheckController(),

		// Article controllers
		getArticlesCtrl:             article.NewGetArticlesController(),
		createArticleCtrl:           article.NewCreateArticleController(),
		getArticlesByTagIDCtrl:      tag.NewGetArticlesByTagIDController(),
		getArticlesByCategoryIDCtrl: category.NewGetArticlesByCategoryIDController(),

		// Category controllers
		getCategoriesCtrl:   category.NewGetCategoriesController(),
		createCategoryCtrl:  category.NewCreateCategoryController(),
		deleteCategoryCtrl:  category.NewDeleteCategoryController(),
		updateCategoryCtrl:  category.NewUpdateCategoryController(),
		getCategoryByIDCtrl: category.NewGetCategoryByIDController(),

		// Tag controllers
		getTagsCtrl:    tag.NewGetTagsController(),
		createTagCtrl:  tag.NewCreateTagController(),
		deleteTagCtrl:  tag.NewDeleteTagController(),
		updateTagCtrl:  tag.NewUpdateTagController(),
		getTagByIDCtrl: tag.NewGetTagByIDController(),
	}
}

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
	// TODO: implement
}

func (s *ServerController) GetPopularArticles(c *gin.Context, params gen.GetPopularArticlesParams) {
	// TODO: implement
}

func (s *ServerController) GetRecentArticles(c *gin.Context, params gen.GetRecentArticlesParams) {
	// TODO: implement
}

func (s *ServerController) DeleteArticleByID(c *gin.Context, articleID string) {
	// TODO: implement
}

func (s *ServerController) UpdateArticleByID(c *gin.Context, articleID string) {
	// TODO: implement
}

func (s *ServerController) PublishArticleByID(c *gin.Context, articleID string) {
	// TODO: implement
}

func (s *ServerController) IncrementArticleViewCount(c *gin.Context, articleID string) {
	// TODO: implement
}

func (s *ServerController) GetArticleByID(c *gin.Context, articleID string) {
	// TODO: implement
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
	// TODO: implement
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
