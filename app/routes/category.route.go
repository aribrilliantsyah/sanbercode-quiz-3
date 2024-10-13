package routes

import (
	"book-management/app/controllers"

	"github.com/gin-gonic/gin"
)

type CategoryRoutes struct {
	categoryController controllers.CategoryController
}

func NewRouteCategory(categoryController controllers.CategoryController) CategoryRoutes {
	return CategoryRoutes{categoryController}
}

func (cr *CategoryRoutes) CategoryRoute(rg *gin.RouterGroup) {

	router := rg.Group("categories")
	router.POST("/", cr.categoryController.CreateCategory)
	router.GET("/", cr.categoryController.GetAllCategories)
	router.PUT("/:id", cr.categoryController.UpdateCategory)
	router.GET("/:id", cr.categoryController.GetCategoryById)
	router.DELETE("/:id", cr.categoryController.DeleteCategoryById)
	router.GET("/:id/books", cr.categoryController.GetBooksByCategoryId)
}
