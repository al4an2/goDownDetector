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
		Name  string `json:"name" binding:"required"`
		Email string `json:"email" binding:"required,email"`
	}

	var new_user user
	if err := c.ShouldBindJSON(&new_user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error parsing JSON: %s", err)})
		return
	}

	lengthUserName := len(new_user.Name)
	if lengthUserName < 1 || lengthUserName > 100 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": fmt.Sprintf(
			"User name: %s with length %d isn't allowed size. Username must be between 1 and 100 characters long.",
			new_user.Name,
			lengthUserName)})
		return
	}

	created_user, err := apiCfg.DB.CreateUser(c, database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      new_user.Name,
		Email:     new_user.Email,
		Usertype:  "user",
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Couldn't create user: %s", err)})
		return
	}

	c.JSON(http.StatusCreated, created_user)
}

func (apiCfg *apiConfig) handlerGetUser(c *gin.Context) {

	userStruct := getUserStruct(c)
	c.JSON(http.StatusOK, userStruct)
}

func (apiCfg *apiConfig) handlerGetAllUsers(c *gin.Context) {

	userStruct := getUserStruct(c)

	if userStruct.Usertype != "admin" {
		c.JSON(http.StatusBadRequest, "Getting user finish with error: You are NOT admin!")
		return
	}

	users, err := apiCfg.DB.GetAllUsers(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("Couldn't get all users: %s", err))
		return
	}
	c.JSON(http.StatusOK, users)
}
