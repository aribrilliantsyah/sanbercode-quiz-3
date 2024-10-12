package controllers

import (
	"context"
	"net/http"

	db "book-management/db/sqlc"
	"book-management/schemas"
	"book-management/util/common"
	"book-management/util/jwt"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	db  *db.Queries
	ctx context.Context
}

func NewAuthController(db *db.Queries, ctx context.Context) *AuthController {
	return &AuthController{db, ctx}
}

// Login godoc
// @Summary Login
// @Description Login with username and password
// @Tags auth
// @Accept json
// @Produce json
// @Param payload body schemas.Login true "Login Data"
// @Success 200 {object} schemas.Response
// @Failure 400 {object} schemas.Response
// @Failure 502 {object} schemas.Response
// @Router /api/auth/login [post]
func (c *AuthController) Login(ctx *gin.Context) {
	var payload *schemas.Login

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	user, err := c.db.GetUserByUsername(ctx, payload.Username)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  "failed",
			"message": "invalid credentials (1)",
		})
		return
	}

	if !jwt.CheckPasswordHash(payload.Password, user.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  "failed",
			"message": "invalid credentials (2)",
		})
		return
	}

	token, err := jwt.GenerateToken(user.ID, user.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"failed":  "failed",
			"message": "failed to generate token",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "successfully logged in",
		"data": gin.H{
			"token": token,
		},
	})
}

// Register godoc
// @Summary Register a new user
// @Description Register a new user with username and password
// @Tags auth
// @Accept json
// @Produce json
// @Param payload body schemas.Register true "Register Data"
// @Success 201 {object} schemas.Response
// @Failure 400 {object} schemas.Response
// @Failure 502 {object} schemas.Response
// @Router /api/auth/register [post]
func (c *AuthController) Register(ctx *gin.Context) {
	var payload *schemas.Register

	// Bind JSON payload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "invalid input data",
		})
		return
	}

	hashedPassword, err := jwt.HashPassword(payload.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": "failed to hash password",
		})
		return
	}

	args := &db.CreateUserParams{
		Username: payload.Username,
		Password: string(hashedPassword),
	}

	user, err := c.db.CreateUser(ctx, *args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": "failed to create user",
			"err":     err.Error(),
		})
		return
	}

	data := schemas.UserResponse{
		ID:         user.ID,
		Username:   user.Username,
		CreatedBy:  common.ConvertNullString(user.CreatedBy),
		CreatedAt:  common.ConvertNullTime(user.CreatedAt),
		ModifiedBy: common.ConvertNullString(user.ModifiedBy),
		ModifiedAt: common.ConvertNullTime(user.ModifiedAt),
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data":   data,
	})
}
