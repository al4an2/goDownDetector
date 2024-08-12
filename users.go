package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/al4an2/goDownDetector/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(c *gin.Context) {
	type user struct {
		Name string `json:"name"`
	}

	var new_user user
	if err := c.ShouldBindJSON(&new_user); err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("Error parsing JSON: %s", err)})
		return
	}
	created_user, err := apiCfg.DB.CreateUser(c, database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      new_user.Name,
	})
	if err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("Couldn't create user: %s", err)})
		return
	}

	c.JSON(http.StatusCreated, created_user)
}
