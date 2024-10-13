package controllers

import (
	"context"
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"book-management/app/schemas"
	db "book-management/db/sqlc"
	"book-management/util/common"
	"book-management/util/jwt"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	db  *db.Queries
	ctx context.Context
}

func NewCategoryController(db *db.Queries, ctx context.Context) *CategoryController {
	return &CategoryController{db, ctx}
}

// CreateCategory godoc
// @Summary Create a new category
// @Description Create a new category with the given payload
// @Tags categories
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token for authorization"
// @Param payload body schemas.CreateCategory true "Category Data"
// @Success 200 {object} schemas.Response
// @Failure 400 {object} schemas.Response
// @Failure 502 {object} schemas.Response
// @Router /api/categories [post]
func (c *CategoryController) CreateCategory(ctx *gin.Context) {
	var payload *schemas.CreateCategory

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	userInfo, err := jwt.GetUserInfo(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}
	username := userInfo.Username

	now := time.Now()
	args := &db.CreateCategoryParams{
		Name:      payload.Name,
		CreatedAt: sql.NullTime{Time: now, Valid: true},
		CreatedBy: sql.NullString{String: username, Valid: true},
	}

	category, err := c.db.CreateCategory(ctx, *args)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	data := schemas.CategoryData{
		ID:         category.ID,
		Name:       category.Name,
		CreatedBy:  common.ConvertNullString(category.CreatedBy),
		CreatedAt:  common.ConvertNullTime(category.CreatedAt),
		ModifiedAt: common.ConvertNullTime(category.ModifiedAt),
		ModifiedBy: common.ConvertNullString(category.ModifiedBy),
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "created successfully",
		"data":    data,
	})
}

// UpdateCategory godoc
// @Summary Update an existing category
// @Description Update a category with the given ID and payload
// @Tags categories
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token for authorization"
// @Param id path int true "Category ID"
// @Param payload body schemas.UpdateCategory true "Category Update Data"
// @Success 200 {object} schemas.Response
// @Failure 400 {object} schemas.Response
// @Failure 404 {object} schemas.Response
// @Failure 502 {object} schemas.Response
// @Router /api/categories/{id} [put]
func (c *CategoryController) UpdateCategory(ctx *gin.Context) {
	var payload *schemas.UpdateCategory
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "invalid category id",
		})
		return
	}

	userInfo, err := jwt.GetUserInfo(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}
	username := userInfo.Username

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	if _, err := c.db.GetCategory(ctx, id); err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{
				"status":  "failed",
				"message": "failed to retrieve category with this id",
			})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	now := time.Now()
	args := &db.UpdateCategoryParams{
		ID:         id,
		Name:       payload.Name,
		ModifiedAt: sql.NullTime{Time: now, Valid: true},
		ModifiedBy: sql.NullString{String: username, Valid: true},
	}

	category, err := c.db.UpdateCategory(ctx, *args)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{
				"status":  "failed",
				"message": "failed to retrieve category with this id",
			})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	data := schemas.CategoryData{
		ID:         category.ID,
		Name:       category.Name,
		CreatedBy:  common.ConvertNullString(category.CreatedBy),
		CreatedAt:  common.ConvertNullTime(category.CreatedAt),
		ModifiedAt: common.ConvertNullTime(category.ModifiedAt),
		ModifiedBy: common.ConvertNullString(category.ModifiedBy),
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "updated successfully",
		"data":    data,
	})
}

// GetCategoryById godoc
// @Summary Get a category by ID
// @Description Retrieve a category by its ID
// @Tags categories
// @Produce json
// @Param Authorization header string true "Bearer token for authorization"
// @Param id path int true "Category ID"
// @Success 200 {object} schemas.Response
// @Failure 400 {object} schemas.Response
// @Failure 404 {object} schemas.Response
// @Failure 502 {object} schemas.Response
// @Router /api/categories/{id} [get]
func (c *CategoryController) GetCategoryById(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "invalid category id",
		})
		return
	}

	category, err := c.db.GetCategory(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{
				"status":  "failed",
				"message": "failed to retrieve category with this id",
			})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	data := schemas.CategoryData{
		ID:         category.ID,
		Name:       category.Name,
		CreatedBy:  common.ConvertNullString(category.CreatedBy),
		CreatedAt:  common.ConvertNullTime(category.CreatedAt),
		ModifiedAt: common.ConvertNullTime(category.ModifiedAt),
		ModifiedBy: common.ConvertNullString(category.ModifiedBy),
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "retrieved successfully",
		"data":    data,
	})
}

// GetAllCategories godoc
// @Summary Get all categories
// @Description Retrieve all categories with pagination
// @Tags categories
// @Produce json
// @Param Authorization header string true "Bearer token for authorization"
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Limit per page" default(10)
// @Success 200 {object} schemas.Response
// @Failure 502 {object} schemas.Response
// @Router /api/categories [get]
func (c *CategoryController) GetAllCategories(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	reqPageID, _ := strconv.Atoi(page)
	reqLimit, _ := strconv.Atoi(limit)
	offset := (reqPageID - 1) * reqLimit

	args := &db.ListCategoriesParams{
		Limit:  int32(reqLimit),
		Offset: int32(offset),
	}

	categories, err := c.db.ListCategories(ctx, *args)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	if categories == nil {
		categories = []db.Category{}
	}

	data := make([]schemas.CategoryData, len(categories))
	for i, category := range categories {
		data[i] = schemas.CategoryData{
			ID:         category.ID,
			Name:       category.Name,
			CreatedBy:  common.ConvertNullString(category.CreatedBy),
			CreatedAt:  common.ConvertNullTime(category.CreatedAt),
			ModifiedAt: common.ConvertNullTime(category.ModifiedAt),
			ModifiedBy: common.ConvertNullString(category.ModifiedBy),
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "retrieved successfully",
		"data":    data,
	})
}

// DeleteCategoryById godoc
// @Summary Delete a category by ID
// @Description Delete a category with the given ID
// @Tags categories
// @Produce json
// @Param Authorization header string true "Bearer token for authorization"
// @Param id path int true "Category ID"
// @Success 204 {object} schemas.Response
// @Failure 400 {object} schemas.Response
// @Failure 404 {object} schemas.Response
// @Failure 502 {object} schemas.Response
// @Router /api/categories/{id} [delete]
func (c *CategoryController) DeleteCategoryById(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "invalid category id",
		})
		return
	}

	if _, err := c.db.GetCategory(ctx, id); err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{
				"status":  "failed",
				"message": "failed to retrieve category with this id",
			})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	err = c.db.DeleteCategory(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "deleted successfully",
	})
}

// GetBooksByCategoryID godoc
// @Summary Get all categories
// @Description Retrieve all categories with pagination
// @Tags categories
// @Produce json
// @Param Authorization header string true "Bearer token for authorization"
// @Param id path int true "Category ID"
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Limit per page" default(10)
// @Success 200 {object} schemas.Response
// @Failure 502 {object} schemas.Response
// @Router /api/categories/{id}/books [get]
func (c *CategoryController) GetBooksByCategoryId(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "invalid category id",
		})
		return
	}

	// Default pagination values
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	reqPageID, _ := strconv.Atoi(page)
	reqLimit, _ := strconv.Atoi(limit)
	offset := (reqPageID - 1) * reqLimit

	// Query parameters for listing books by category
	args := &db.GetBooksByCategoryParams{
		CategoryID: id,
		Limit:      int32(reqLimit),
		Offset:     int32(offset),
	}

	books, err := c.db.GetBooksByCategory(ctx, *args)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	if books == nil {
		books = []db.Book{}
	}

	// Prepare response data
	data := make([]schemas.BookData, len(books))
	for i, book := range books {
		data[i] = schemas.BookData{
			ID:          book.ID,
			Title:       book.Title,
			Description: common.ConvertNullString(book.Description),
			ImageUrl:    common.ConvertNullString(book.ImageUrl),
			ReleaseYear: common.ConvertNullInt32(book.ReleaseYear),
			Price:       common.ConvertNullInt32(book.Price),
			TotalPage:   common.ConvertNullInt32(book.TotalPage),
			Thickness:   common.ConvertNullString(book.Thickness),
			CategoryID:  book.CategoryID,
			CreatedBy:   common.ConvertNullString(book.CreatedBy),
			CreatedAt:   common.ConvertNullTime(book.CreatedAt),
			ModifiedAt:  common.ConvertNullTime(book.ModifiedAt),
			ModifiedBy:  common.ConvertNullString(book.ModifiedBy),
		}
	}

	// Send response
	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "retrieved successfully",
		"data":    data,
	})
}
