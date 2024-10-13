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

type BookController struct {
	db  *db.Queries
	ctx context.Context
}

func NewBookController(db *db.Queries, ctx context.Context) *BookController {
	return &BookController{db, ctx}
}

// CreateBook godoc
// @Summary Create a new book
// @Description Create a new book with the given payload
// @Tags books
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token for authorization"
// @Param payload body schemas.CreateBook true "Book Data"
// @Success 200 {object} schemas.Response
// @Failure 400 {object} schemas.Response
// @Failure 502 {object} schemas.Response
// @Router /api/books [post]
func (b *BookController) CreateBook(ctx *gin.Context) {
	var payload *schemas.CreateBook

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	if _, err := b.db.GetCategory(ctx, payload.CategoryID); err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{
				"status":  "failed",
				"message": "failed to retrieve category with this category id",
			})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	// Validate release_year
	if payload.ReleaseYear < 1980 || payload.ReleaseYear > 2024 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "release_year must be between 1980 and 2024",
		})
		return
	}

	// Determine thickness based on total page
	var thickness string
	if payload.TotalPage > 100 {
		thickness = "tebal"
	} else {
		thickness = "tipis"
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
	args := &db.CreateBookParams{
		Title:       payload.Title,
		Description: sql.NullString{String: payload.Description, Valid: payload.Description != ""},
		ImageUrl:    sql.NullString{String: payload.ImageUrl, Valid: payload.ImageUrl != ""},
		ReleaseYear: sql.NullInt32{Int32: payload.ReleaseYear, Valid: payload.ReleaseYear != 0},
		Price:       sql.NullInt32{Int32: payload.Price, Valid: payload.Price != 0},
		TotalPage:   sql.NullInt32{Int32: payload.TotalPage, Valid: payload.TotalPage != 0},
		Thickness:   sql.NullString{String: thickness, Valid: true}, // Use calculated thickness
		CategoryID:  payload.CategoryID,
		CreatedAt:   sql.NullTime{Time: now, Valid: true},
		CreatedBy:   sql.NullString{String: username, Valid: true},
	}

	book, err := b.db.CreateBook(ctx, *args)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	data := schemas.BookData{
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

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "created successfully",
		"data":    data,
	})
}

// UpdateBook godoc
// @Summary Update an existing book
// @Description Update a book with the given ID and payload
// @Tags books
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token for authorization"
// @Param id path int true "Book ID"
// @Param payload body schemas.UpdateBook true "Book Update Data"
// @Success 200 {object} schemas.Response
// @Failure 400 {object} schemas.Response
// @Failure 404 {object} schemas.Response
// @Failure 502 {object} schemas.Response
// @Router /api/books/{id} [put]
func (b *BookController) UpdateBook(ctx *gin.Context) {
	var payload *schemas.UpdateBook
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "invalid book id",
		})
		return
	}

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	if _, err := b.db.GetCategory(ctx, payload.CategoryID); err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{
				"status":  "failed",
				"message": "failed to retrieve category with this category id",
			})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	// Validate release_year
	if payload.ReleaseYear < 1980 || payload.ReleaseYear > 2024 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "release_year must be between 1980 and 2024",
		})
		return
	}

	// Determine thickness based on total page
	var thickness string
	if payload.TotalPage > 100 {
		thickness = "tebal"
	} else {
		thickness = "tipis"
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
	args := &db.UpdateBookParams{
		ID:          id,
		Title:       payload.Title,
		Description: sql.NullString{String: payload.Description, Valid: payload.Description != ""},
		ImageUrl:    sql.NullString{String: payload.ImageUrl, Valid: payload.ImageUrl != ""},
		ReleaseYear: sql.NullInt32{Int32: payload.ReleaseYear, Valid: payload.ReleaseYear != 0},
		Price:       sql.NullInt32{Int32: payload.Price, Valid: payload.Price != 0},
		TotalPage:   sql.NullInt32{Int32: payload.TotalPage, Valid: payload.TotalPage != 0},
		Thickness:   sql.NullString{String: thickness, Valid: true}, // Use calculated thickness
		CategoryID:  payload.CategoryID,
		ModifiedAt:  sql.NullTime{Time: now, Valid: true},
		ModifiedBy:  sql.NullString{String: username, Valid: true},
	}

	book, err := b.db.UpdateBook(ctx, *args)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{
				"status":  "failed",
				"message": "failed to retrieve book with this id",
			})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	data := schemas.BookData{
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

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "updated successfully",
		"data":    data,
	})
}

// GetBookById godoc
// @Summary Get a book by ID
// @Description Retrieve a book by its ID
// @Tags books
// @Produce json
// @Param Authorization header string true "Bearer token for authorization"
// @Param id path int true "Book ID"
// @Success 200 {object} schemas.Response
// @Failure 400 {object} schemas.Response
// @Failure 404 {object} schemas.Response
// @Failure 502 {object} schemas.Response
// @Router /api/books/{id} [get]
func (b *BookController) GetBookById(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "invalid book id",
		})
		return
	}

	book, err := b.db.GetBook(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{
				"status":  "failed",
				"message": "failed to retrieve book with this id",
			})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	data := schemas.BookData{
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

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "retrieved successfully",
		"data":    data,
	})
}

// GetAllBooks godoc
// @Summary Get all books
// @Description Retrieve all books with pagination
// @Tags books
// @Produce json
// @Param Authorization header string true "Bearer token for authorization"
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Limit per page" default(10)
// @Success 200 {object} schemas.Response
// @Failure 502 {object} schemas.Response
// @Router /api/books [get]
func (b *BookController) GetAllBooks(ctx *gin.Context) {
	// Default pagination values
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	reqPageID, _ := strconv.Atoi(page)
	reqLimit, _ := strconv.Atoi(limit)
	offset := (reqPageID - 1) * reqLimit

	// Query parameters for listing books
	args := &db.ListBooksParams{
		Limit:  int32(reqLimit),
		Offset: int32(offset),
	}

	books, err := b.db.ListBooks(ctx, *args)
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

// DeleteBookById godoc
// @Summary Delete a book by ID
// @Description Delete a book with the given ID
// @Tags books
// @Param Authorization header string true "Bearer token for authorization"
// @Param id path int true "Book ID"
// @Success 200 {object} schemas.Response
// @Failure 400 {object} schemas.Response
// @Failure 404 {object} schemas.Response
// @Failure 502 {object} schemas.Response
// @Router /api/books/{id} [delete]
func (b *BookController) DeleteBookById(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "invalid book id",
		})
		return
	}

	err = b.db.DeleteBook(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{
				"status":  "failed",
				"message": "book not found",
			})
			return
		}
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
