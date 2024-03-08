package middleware

import (
	"equipment-management/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication(c *gin.Context) {
	apiKey := c.Request.Header.Get("x-api-key")
	if apiKey == "" {
		apiKey = fromCookies(c)
		if apiKey == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Api key is empty",
			})
			return
		}
	}

	if apiKey != config.ServerConfig().ApiKey {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Api key is invalid",
		})
		return
	}

	c.Next()
}

func fromCookies(c *gin.Context) string {
	v, err := c.Cookie("api_key")
	if err != nil {
		return ""
	}

	return v
}
