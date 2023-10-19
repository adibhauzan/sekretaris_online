package controllers

import (
	"net/http"
	"time"

	"github.com/adibhauzan/sekretaris_online_backend/models"
	"github.com/adibhauzan/sekretaris_online_backend/repository"
	"github.com/adibhauzan/sekretaris_online_backend/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var InvalidatedTokens = make(map[string]bool)

type UserController struct {
	UserRepo repository.UserRepository
}

func NewUserController(repo repository.UserRepository) *UserController {
	return &UserController{UserRepo: repo}
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user.Password = hashedPassword

	if err := c.UserRepo.CreateUser(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

func (c *UserController) UserLogin(ctx *gin.Context) {
	var request models.User
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := c.UserRepo.GetUserByUsername(request.Username)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if err := utils.VerifyPassword(user.Password, request.Password); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	secretKey := "rahasia"
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func (c *UserController) UserLogout(ctx *gin.Context) {
	tokenString := ctx.GetHeader("Authorization")

	if InvalidatedTokens[tokenString] {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid"})
		return
	}
	if tokenString == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid"})
		return
	}

	InvalidatedTokens[tokenString] = true

	ctx.JSON(http.StatusOK, gin.H{"message": "Logout berhasil"})
}
