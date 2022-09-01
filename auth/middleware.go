package auth

import (
	"context"
	"crowdfund-go/helper"
	"crowdfund-go/user"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var userCtxKey = &contextKey{"currentUser"}

type contextKey struct {
	name string
}

func Middleware(authService Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")

		if !strings.Contains(header, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenHeader := ""
		spltdHeader := strings.Split(header, " ")
		if len(spltdHeader) == 2 {
			tokenHeader = spltdHeader[1]
		}

		token, err := authService.ValidateToken(tokenHeader)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)

		ctx := context.WithValue(c.Request.Context(), userCtxKey, user)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func ForContext(ctx context.Context) user.User {
	raw, _ := ctx.Value(userCtxKey).(user.User)
	return raw
}
