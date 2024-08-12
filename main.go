package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/ready", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ready",
		})
	})

	//cheking site-status func
	router.GET("/site", func(c *gin.Context) {
		log.Println("The URL:", c.Request.Host+c.Request.URL.Path)
		link := c.Query("link")
		log.Println(link)
		response, err := http.Get(link)
		if err != nil {
			log.Println("Error!!!!: ", err)
		} else {
			log.Println(response.StatusCode)
		}

	})

	router.Run(":8080")

	server := &http.Server{
		Handler: router,
		Addr:    ":8080",
	}

	log.Println("Server starting on port :8080")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
