package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/al4an2/goDownDetector/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateSiteFollow(c *gin.Context) {
	type siteFollow struct {
		SiteID uuid.UUID `json:"site_id" binding:"required"`
	}

	userStruct := getUserStruct(c)

	var newSiteFollow siteFollow
	if err := c.ShouldBindJSON(&newSiteFollow); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error parsing JSON: %s", err)})
		return
	}

	createdSiteFollow, err := apiCfg.DB.CreateSiteFollow(c, database.CreateSiteFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    userStruct.ID,
		SiteID:    newSiteFollow.SiteID,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Couldn't create new follow for site: %s", err)})
		return
	}

	c.JSON(http.StatusCreated, createdSiteFollow)
}

func (apiCfg *apiConfig) handlerGetSiteFollows(c *gin.Context) {

	sites, err := apiCfg.DB.GetSiteFollows(c.Request.Context(), getUserStruct(c).ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Couldn't get you site follows: %s", err)})
		return
	}
	c.JSON(http.StatusOK, sites)
}

func (apiCfg *apiConfig) handlerGetAllSiteFollows(c *gin.Context) {

	usertype := getUserStruct(c).Usertype
	if usertype != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Getting user finish with error: You are NOT admin!"})
		return
	}

	sites, err := apiCfg.DB.GetAllSiteFollows(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Haven't any site follows in database or some error on DB request: %s", err)})
		return
	}
	c.JSON(http.StatusOK, sites)
}
