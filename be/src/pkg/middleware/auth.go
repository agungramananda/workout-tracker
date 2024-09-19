package middleware

import (
	"net/http"
	"workout-tracker/m/v0.0.0/src/pkg/utils"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func (c *gin.Context){
		err := utils.VerifyAccessToken(c)
		if err != nil {
			errResponse := utils.NewErrorResponse(http.StatusUnauthorized, "Unauthorized", "Invalid access token")
			c.JSON(http.StatusUnauthorized, errResponse)
			c.Abort()
			return
		}
		c.Next()
	}
}