package middleware

import (
	"net/http"
	"github.com/adibhauzan/sekretaris_online_backend/controllers"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        token := ctx.GetHeader("Authorization")

        if token == "" {
            ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid"})
            ctx.Abort()
            return
        }

        if controllers.InvalidatedTokens[token] {
            ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid"})
            ctx.Abort()
            return
        }

        ctx.Next()
    }
}


func IsUserLoggedIn(ctx *gin.Context) bool {
	
	token := ctx.GetHeader("Authorization")
	return token != ""
}
