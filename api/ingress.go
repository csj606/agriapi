package api

import (
	"context"
	"log"
	"os"

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

func getPlantingSeason(c *gin.Context) {
	response := queryPlantingSeason()
	c.IndentedJSON(200, response)
}

func setUpServer() *gin.Engine {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	context.Background()
	router := gin.Default()
	_, err = os.Stat("agriapi.db")
	if os.IsNotExist(err) {
		CreateDatabase()
	}

	router.GET("/in-season", getInSeason)
	router.GET("/planting-periods", getPlantingSeason)
	router.GET("/ping", online)
	return router
}

func main() {
	router := setUpServer()
	router.Run("localhost:9876")
}
