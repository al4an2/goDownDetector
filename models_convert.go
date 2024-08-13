package main

import (
	"fmt"
	"net/http"

	"github.com/al4an2/goDownDetector/internal/database"
	"github.com/gin-gonic/gin"
)

func getUserStruct(c *gin.Context) *database.User {

	user, exist := c.Get("user")
	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found in query Context"})
		c.Abort()
		return nil
	}

	fmt.Println(user)

	userStruct, ok := user.(*database.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cast user from query Context"})
		c.Abort()
		return nil
	}
	return userStruct
}
