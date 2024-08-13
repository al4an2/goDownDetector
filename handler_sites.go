package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/al4an2/goDownDetector/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateSite(c *gin.Context) {
	type site struct {
		Name string `json:"name" binding:"required"`
		Url  string `json:"url" binding:"required"`
	}

	userStruct := getUserStruct(c)

	var new_site site
	if err := c.ShouldBindJSON(&new_site); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error parsing JSON: %s", err)})
		return
	}

	created_site, err := apiCfg.DB.CreateSite(c, database.CreateSiteParams{
		ID:          uuid.New(),
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
		Name:        new_site.Name,
		Url:         new_site.Url,
		AddedByUser: userStruct.ID,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Couldn't added new site to system: %s", err)})
		return
	}

	c.JSON(http.StatusCreated, created_site)
}
