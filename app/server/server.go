package server

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"book-management/app/controllers"
	"book-management/app/middlewares"
	"book-management/app/routes"
	dbCon "book-management/db/sqlc"
	"book-management/util/config"
	"book-management/util/swagger"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Server struct {
	engine             *gin.Engine
	db                 *dbCon.Queries
	ctx                context.Context
	authController     controllers.AuthController
	categoryController controllers.CategoryController
	bookController     controllers.BookController
	authRoutes         routes.AuthRoutes
	categoryRoutes     routes.CategoryRoutes
	bookRoutes         routes.BookRoutes
}

func NewServer(config config.Config) *Server {
	ctx := context.TODO()

	conn, err := sql.Open(config.DbDriver, config.DbSource)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	db := dbCon.New(conn)
	fmt.Println("PostgreSql connected successfully...")

	// Initialize controllers and routes
	authController := *controllers.NewAuthController(db, ctx)
	categoryController := *controllers.NewCategoryController(db, ctx)
	bookController := *controllers.NewBookController(db, ctx)

	authRoutes := routes.NewRouteAuth(authController)
	categoryRoutes := routes.NewRouteCategory(categoryController)
	bookRoutes := routes.NewRouteBook(bookController)

	server := &Server{
		engine:             gin.Default(),
		db:                 db,
		ctx:                ctx,
		authController:     authController,
		categoryController: categoryController,
		bookController:     bookController,
		authRoutes:         authRoutes,
		categoryRoutes:     categoryRoutes,
		bookRoutes:         bookRoutes,
	}

	// Initialize Swagger
	swagger.Initialize(server.engine)

	// Set up routes
	server.setupRoutes()
	return server
}

func (s *Server) setupRoutes() {
	s.engine.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "welcome to api, docs: /docs/index.html",
		})
	})

	// Public routes
	group := s.engine.Group("/api")
	s.authRoutes.AuthRoute(group)

	// Protected routes with AuthMiddleware
	protected := s.engine.Group("/api")
	protected.Use(middlewares.AuthMiddleware())
	s.categoryRoutes.CategoryRoute(protected)
	s.bookRoutes.BookRoute(protected)

	// Handle 404
	s.engine.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "failed",
			"message": fmt.Sprintf("The specified route %s not found", ctx.Request.URL),
		})
	})
}

func (s *Server) Run() error {
	config, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %v", err)
	}
	return s.engine.Run(":" + config.ServerAddress)
}
