package main

import (
	"fmt"
	"net/http"

	"github.com/al4an2/goDownDetector/internal/auth"
	"github.com/gin-gonic/gin"
)

func (apiCfg *apiConfig) middlewareAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey, err := auth.GetAPIKey(c.Request.Header.Get("Authorization"))
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": fmt.Sprintf("Auth apiKey is error: %s", err)})
			c.Abort()
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(c.Request.Context(), apiKey)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Getting user finished with error: %s", err)})
			c.Abort()
			return
		}

		// Store the user in the context
		c.Set("user", &user)

		// Continue to the next middleware/handler
		c.Next()
	}
}
