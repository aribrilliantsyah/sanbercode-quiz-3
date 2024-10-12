package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"book-management/controllers"
	dbCon "book-management/db/sqlc"
	"book-management/middlewares"
	"book-management/routes"
	"book-management/util"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var (
	server *gin.Engine
	db     *dbCon.Queries
	ctx    context.Context

	AuthController controllers.AuthController
	AuthRoutes     routes.AuthRoutes

	CategoryController controllers.CategoryController
	CategoryRoutes     routes.CategoryRoutes
)

func init() {
	ctx = context.TODO()
	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatalf("could not loadutil: %v", err)
	}

	conn, err := sql.Open(config.DbDriver, config.DbSource)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	db = dbCon.New(conn)

	fmt.Println("PostgreSql connected successfully...")

	CategoryController = *controllers.NewCategoryController(db, ctx)
	CategoryRoutes = routes.NewRouteCategory(CategoryController)

	AuthController = *controllers.NewAuthController(db, ctx)
	AuthRoutes = routes.NewRouteAuth(AuthController)

	server = gin.Default()
}

func main() {
	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	util.InitSwagger(server)

	router := server
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "welcome to api, docs: /docs",
		})
	})
	group := router.Group("/api")
	AuthRoutes.AuthRoute(group)

	protected := router.Group("/api")
	protected.Use(middlewares.AuthMiddleware())
	CategoryRoutes.CategoryRoute(protected)

	server.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "failed", "message": fmt.Sprintf("The specified route %s not found", ctx.Request.URL)})
	})

	log.Fatal(server.Run(":" + config.ServerAddress))
}
