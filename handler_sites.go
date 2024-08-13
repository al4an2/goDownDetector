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

func (apiCfg *apiConfig) handlerGetSites(c *gin.Context) {

	sites, err := apiCfg.DB.GetSites(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Couldn't get site's list: %s", err)})
		return
	}
	c.JSON(http.StatusOK, sites)
}

func (apiCfg *apiConfig) handlerGetMyAddedSites(c *gin.Context) {

	userStruct := getUserStruct(c)

	sites, err := apiCfg.DB.GetMyAddedSites(c.Request.Context(), userStruct.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("The sites you added couldn't be found: %s", err)})
		return
	}

	if sites == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The sites you added couldn't be found: You haven't added sites to the system yet. This can be done using the /sites POST request."})
		return
	}
	c.JSON(http.StatusOK, sites)
}

func (apiCfg *apiConfig) handlerGetAllSitesInfo(c *gin.Context) {

	usertype := getUserStruct(c).Usertype
	if usertype != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Getting user finish with error: You are NOT admin!"})
		return
	}

	sites, err := apiCfg.DB.GetAllSitesInfo(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Haven't any site in list or some error on DB request: %s", err)})
		return
	}
	c.JSON(http.StatusOK, sites)
}
