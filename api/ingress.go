package api

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func getInSeason(c *gin.Context) {
	response := queryInSeason()
	c.IndentedJSON(200, response)
}

func online(c *gin.Context) {
	c.IndentedJSON(200, gin.H{"message": "received"})
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	context.Background()

	router := gin.Default()
	router.GET("/in-season", getInSeason)
	router.GET("/ping", online)
	router.Run("localhost:9876")
}
