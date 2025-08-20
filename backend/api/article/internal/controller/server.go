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
	getArticlesCtrl               *article.GetArticlesController
	createArticleCtrl             *article.CreateArticleController
	getArticlesByTagIdCtrl        *tag.GetArticlesByTagIDController
	getArticlesByCategoryIdCtrl   *category.GetArticlesByCategoryIDController

	// Health controller
	healthcheckCtrl *health.HealthcheckController

	// Category controllers
	getCategoriesCtrl     *category.GetCategoriesController
	createCategoryCtrl    *category.CreateCategoryController
	deleteCategoryCtrl    *category.DeleteCategoryController
	updateCategoryCtrl    *category.UpdateCategoryController
	getCategoryByIdCtrl   *category.GetCategoryByIDController

	// Tag controllers
	getTagsCtrl      *tag.GetTagsController
	createTagCtrl    *tag.CreateTagController
	deleteTagCtrl    *tag.DeleteTagController
	updateTagCtrl    *tag.UpdateTagController
	getTagByIdCtrl   *tag.GetTagByIDController
}

// NewServerController creates a new server controller
func NewServerController() *ServerController {
	return &ServerController{
		// Health controller
		healthcheckCtrl: health.NewHealthcheckController(),

		// Article controllers
		getArticlesCtrl:               article.NewGetArticlesController(),
		createArticleCtrl:             article.NewCreateArticleController(),
		getArticlesByTagIdCtrl:        tag.NewGetArticlesByTagIDController(),
		getArticlesByCategoryIdCtrl:   category.NewGetArticlesByCategoryIDController(),

		// Category controllers
		getCategoriesCtrl:     category.NewGetCategoriesController(),
		createCategoryCtrl:    category.NewCreateCategoryController(),
		deleteCategoryCtrl:    category.NewDeleteCategoryController(),
		updateCategoryCtrl:    category.NewUpdateCategoryController(),
		getCategoryByIdCtrl:   category.NewGetCategoryByIDController(),

		// Tag controllers
		getTagsCtrl:      tag.NewGetTagsController(),
		createTagCtrl:    tag.NewCreateTagController(),
		deleteTagCtrl:    tag.NewDeleteTagController(),
		updateTagCtrl:    tag.NewUpdateTagController(),
		getTagByIdCtrl:   tag.NewGetTagByIDController(),
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

// Stub implementations for other endpoints (TODO: implement)
func (s *ServerController) GetArticlesByMonth(c *gin.Context, year gen.YearParam, month gen.MonthParam, params gen.GetArticlesByMonthParams) {
	// TODO: implement
}

func (s *ServerController) GetPopularArticles(c *gin.Context, params gen.GetPopularArticlesParams) {
	// TODO: implement
}

func (s *ServerController) GetRecentArticles(c *gin.Context, params gen.GetRecentArticlesParams) {
	// TODO: implement
}

func (s *ServerController) DeleteArticleById(c *gin.Context, articleId string) {
	// TODO: implement
}

func (s *ServerController) UpdateArticleById(c *gin.Context, articleId string) {
	// TODO: implement
}

func (s *ServerController) PublishArticleById(c *gin.Context, articleId string) {
	// TODO: implement
}

func (s *ServerController) IncrementArticleViewCount(c *gin.Context, articleId string) {
	// TODO: implement
}

func (s *ServerController) GetArticleById(c *gin.Context, articleId string) {
	// TODO: implement
}

func (s *ServerController) GetCategories(c *gin.Context) {
	s.getCategoriesCtrl.GetCategories(c)
}

func (s *ServerController) CreateCategory(c *gin.Context) {
	s.createCategoryCtrl.CreateCategory(c)
}

func (s *ServerController) DeleteCategoryById(c *gin.Context, categoryId string) {
	s.deleteCategoryCtrl.Handle(c, categoryId)
}

func (s *ServerController) UpdateCategoryById(c *gin.Context, categoryId string) {
	var req gen.CategoryUpdate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	s.updateCategoryCtrl.Handle(c, categoryId, req)
}

func (s *ServerController) GetCategoryById(c *gin.Context, categoryId string) {
	s.getCategoryByIdCtrl.Handle(c, categoryId)
}

func (s *ServerController) GetArticlesByCategoryId(c *gin.Context, categoryId string, params gen.GetArticlesByCategoryIdParams) {
	s.getArticlesByCategoryIdCtrl.Handle(c, categoryId, params)
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

func (s *ServerController) DeleteTagById(c *gin.Context, tagId string) {
	s.deleteTagCtrl.Handle(c, tagId)
}

func (s *ServerController) UpdateTagById(c *gin.Context, tagId string) {
	var req gen.TagUpdate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	s.updateTagCtrl.Handle(c, tagId, req)
}

func (s *ServerController) GetTagById(c *gin.Context, tagId string) {
	s.getTagByIdCtrl.Handle(c, tagId)
}

func (s *ServerController) GetArticlesByTagId(c *gin.Context, tagId string, params gen.GetArticlesByTagIdParams) {
	s.getArticlesByTagIdCtrl.Handle(c, tagId, params)
}
