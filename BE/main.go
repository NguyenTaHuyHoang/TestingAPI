package main

import (
	"log"
	"wan-api-kol-event/Controllers"
	"wan-api-kol-event/Initializers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	Initializers.LoadEnvironmentVariables()
	Initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	// Use the CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://testing-api-lac.vercel.app"}, // Allow your frontend's origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin*", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Define your Gin routes here
	r.GET("/kols", Controllers.GetKolsController(Initializers.DB))

	// Run Gin server
	if err := r.Run(":8081"); err != nil {
		log.Println("Failed to start server")
	}
}
