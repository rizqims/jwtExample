package middleware

import (
	"apilaundry/service"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware interface {
	CheckToken(roles ...string) gin.HandlerFunc
}

type authMiddleware struct {
	jwtService service.JwtService
}

type payloadHeader struct{
	AuthorizationHeader string `header:"Authorization"`
}

func (a *authMiddleware) CheckToken(roles ...string) gin.HandlerFunc{
	return func( c *gin.Context){
		header := c.GetHeader("Authorization")
		token := strings.Replace(header, "Bearer ","",-1)
		claims, err := a.jwtService.VerifyToken(token)
		if err != nil {
			fmt.Println(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message":"Unauthorized"})
			return
		}
		c.Set("userId", claims["userId"])

		var validRole bool
		for _, r := range roles{
			if r == claims["role"]{
				validRole = true
				break
			}
		}
		if !validRole{
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message":"Forbidden Access"})
			return
		}
		c.Next()
	}
}

func NewAuthMiddleware(jwtService service.JwtService) AuthMiddleware{
	return &authMiddleware{jwtService: jwtService}
}