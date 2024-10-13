package routes

import (
	"book-management/app/controllers"

	"github.com/gin-gonic/gin"
)

type BookRoutes struct {
	bookController controllers.BookController
}

func NewRouteBook(bookController controllers.BookController) BookRoutes {
	return BookRoutes{bookController}
}

func (cr *BookRoutes) BookRoute(rg *gin.RouterGroup) {

	router := rg.Group("books")
	router.POST("/", cr.bookController.CreateBook)
	router.GET("/", cr.bookController.GetAllBooks)
	router.PUT("/:id", cr.bookController.UpdateBook)
	router.GET("/:id", cr.bookController.GetBookById)
	router.DELETE("/:id", cr.bookController.DeleteBookById)
}
