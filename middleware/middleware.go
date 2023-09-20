package middleware

import (
	"github.com/CloudGod5/ecommerce-go/tokens"
	"net/http"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return function(c *gin.Context) {
		ClientToken := c.Request.Header.Get("token")
		if ClientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error":"No Authorization header provided"})
			c.Abort()
			return
		}
		Claims, err := token,ValidateToken(ClientToken)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
			c.Abort()
			return
		}

		c.Set("email", Claims.Email)
		c.Set("uid", Claims.Uid)
		c.Next()
	}
}