package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yousefzinsazk78/simple_note_api/internal/database"
)

func JWTAuth(userStorer database.UserStorer) gin.HandlerFunc {
	return func(context *gin.Context) {
		err := ValidateJWT(context)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
			context.Abort()
			return
		}
		error := ValidateAdminRoleJWT(context)
		if error != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Only Administrator is allowed to perform this action"})
			context.Abort()
			return
		}

		user := CurrentUser(context, userStorer)
		context.Set("userModel", user)

		context.Next()
	}
}

func JWTAuthAuthor(userStore database.UserStorer) gin.HandlerFunc {
	return func(context *gin.Context) {
		err := ValidateJWT(context)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
			context.Abort()
			return
		}
		error := ValidateAuthorRoleJWT(context)
		if error != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Only registered Customers are allowed to perform this action"})
			context.Abort()
			return
		}

		user := CurrentUser(context, userStore)

		context.Set("userModel", user)

		context.Next()
	}
}
