package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/al4an2/goDownDetector/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/sqlc-dev/sqlc"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {

	//start initialization
	godotenv.Load(".env")
	portString := os.Getenv("PORT")
	admin_login := os.Getenv("admin_login")
	admin_email := os.Getenv("admin_email")

	router := gin.Default()

	dbUrl := os.Getenv("DB_url")
	if dbUrl == "" {
		log.Fatal("Database url ('DB_url') is not found in the environmental")
	}

	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("Can't connect to database:", err)
	}

	db := database.New(conn)
	apiCfg := apiConfig{
		DB: db,
	}

	//create admin
	apiCfg.createAdmin(admin_login, admin_email, context.TODO())
	//routing
	router.GET("/ready", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ready",
		})
	})

	router.GET("/error", func(c *gin.Context) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
		})
	})

	router.POST("/users", func(c *gin.Context) {
		apiCfg.handlerCreateUser(c)
	})
	router.GET("/users", func(c *gin.Context) {
		apiCfg.handlerGetUser(c)
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

	server := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Println("Server starting on port :8080")
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
