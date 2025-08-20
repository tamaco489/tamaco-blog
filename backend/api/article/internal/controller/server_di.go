// Package controller provides HTTP request handlers and routing logic for the article API.
// This package implements the ServerInterface from the generated OpenAPI specification
// and orchestrates the interaction between HTTP requests and business logic through controllers.
package controller

import (
	"context"

	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/controller/article"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/controller/category"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/controller/health"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/controller/search"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/controller/tag"
	"github.com/tamaco489/tamaco-blog/backend/api/article/internal/library/config"
)

type ServerController struct {
	healthcheckCtrl *health.HealthcheckController

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

	getCategoriesCtrl   *category.GetCategoriesController
	createCategoryCtrl  *category.CreateCategoryController
	deleteCategoryCtrl  *category.DeleteCategoryController
	updateCategoryCtrl  *category.UpdateCategoryController
	getCategoryByIDCtrl *category.GetCategoryByIDController

	getTagsCtrl    *tag.GetTagsController
	createTagCtrl  *tag.CreateTagController
	deleteTagCtrl  *tag.DeleteTagController
	updateTagCtrl  *tag.UpdateTagController
	getTagByIDCtrl *tag.GetTagByIDController

	searchArticlesCtrl *search.SearchArticlesController
}

func NewServerController(ctx context.Context, cfg *config.Config) *ServerController {
	return &ServerController{
		healthcheckCtrl: health.NewHealthcheckController(),

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

		getCategoriesCtrl:   category.NewGetCategoriesController(),
		createCategoryCtrl:  category.NewCreateCategoryController(),
		deleteCategoryCtrl:  category.NewDeleteCategoryController(),
		updateCategoryCtrl:  category.NewUpdateCategoryController(),
		getCategoryByIDCtrl: category.NewGetCategoryByIDController(),

		getTagsCtrl:    tag.NewGetTagsController(ctx, cfg),
		createTagCtrl:  tag.NewCreateTagController(),
		deleteTagCtrl:  tag.NewDeleteTagController(),
		updateTagCtrl:  tag.NewUpdateTagController(),
		getTagByIDCtrl: tag.NewGetTagByIDController(),

		searchArticlesCtrl: search.NewSearchArticlesController(),
	}
}
