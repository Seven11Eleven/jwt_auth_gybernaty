package middleware

import (
	"net/http"
	"strings"

	"github.com/Seven11Eleven/jwt_auth_gybernaty/internal/utils"
	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware(secret string) gin.HandlerFunc{
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		tkn := strings.Split(authHeader, " ")
		if len(tkn) == 2 {
			authToken := tkn[1]
			isAuthorized, err := utils.IsAuthorized(authToken)
			if isAuthorized{
				authorID, err := utils.ExtractIDFromToken(authToken)
				if err != nil{
					c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
					c.Abort()
					return
				}
				c.Set("authorID", authorID.String())
				c.Next()
				return
			}
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized!"})
		c.Abort()
	}
}