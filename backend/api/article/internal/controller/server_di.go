// Package controller provides HTTP controllers for the API.
package controller

import (
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/controller/article"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/controller/category"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/controller/health"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/controller/search"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/controller/tag"
)

// ServerController aggregates all controllers and implements gen.ServerInterface
type ServerController struct {
	// Article controllers
	getArticlesCtrl             *article.GetArticlesController
	createArticleCtrl           *article.CreateArticleController
	getArticlesByMonthCtrl      *article.GetArticlesByMonthController
	getPopularArticlesCtrl      *article.GetPopularArticlesController
	getRecentArticlesCtrl       *article.GetRecentArticlesController
	deleteArticleCtrl           *article.DeleteArticleController
	updateArticleByIDCtrl       *article.UpdateArticleByIDController
	publishArticleByIDCtrl      *article.PublishArticleByIDController
	incrementViewCountCtrl      *article.IncrementArticleViewCountController
	getArticleByIDCtrl          *article.GetArticleByIDController
	getArticlesByTagIDCtrl      *tag.GetArticlesByTagIDController
	getArticlesByCategoryIDCtrl *category.GetArticlesByCategoryIDController

	// Search controllers
	searchArticlesCtrl *search.SearchArticlesController

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
		getArticlesByMonthCtrl:      article.NewGetArticlesByMonthController(),
		getPopularArticlesCtrl:      article.NewGetPopularArticlesController(),
		getRecentArticlesCtrl:       article.NewGetRecentArticlesController(),
		deleteArticleCtrl:           article.NewDeleteArticleController(),
		updateArticleByIDCtrl:       article.NewUpdateArticleByIDController(),
		publishArticleByIDCtrl:      article.NewPublishArticleByIDController(),
		incrementViewCountCtrl:      article.NewIncrementArticleViewCountController(),
		getArticleByIDCtrl:          article.NewGetArticleByIDController(),
		getArticlesByTagIDCtrl:      tag.NewGetArticlesByTagIDController(),
		getArticlesByCategoryIDCtrl: category.NewGetArticlesByCategoryIDController(),

		// Search controllers
		searchArticlesCtrl: search.NewSearchArticlesController(),

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
