package api

import (
	"fp-super-bootcamp-go/config"
	"fp-super-bootcamp-go/docs"
	"fp-super-bootcamp-go/routes"
	"fp-super-bootcamp-go/utils"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	App *gin.Engine
)

func init() {
	App = gin.New()
	// for load godotenv
	// for env
	environment := utils.Getenv("ENVIRONMENT", "production")

	if environment == "development" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	//programmatically set swagger info
	docs.SwaggerInfo.Title = "Review Book REST API"
	docs.SwaggerInfo.Description = "This is REST API Review Book."
	docs.SwaggerInfo.Version = "1.0"
	if environment == "development" {
		docs.SwaggerInfo.Host = utils.Getenv("HOST", "localhost:8080")
		docs.SwaggerInfo.Schemes = []string{"http", "https"}
	} else {
		docs.SwaggerInfo.Host = os.Getenv("HOST")
		docs.SwaggerInfo.Schemes = []string{"https"}
	}

	db := config.ConnectDataBase()

	routes.SetupRouter(db, App)
}

// Entrypoint
func Handler(w http.ResponseWriter, r *http.Request) {
	App.ServeHTTP(w, r)
}
