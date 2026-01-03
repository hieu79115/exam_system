package main

import (
	"exam-system/config"
	"exam-system/internal/app"
	"exam-system/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.InitDB()

	myHandlers := app.InitApp(db)

	r := gin.Default()
	routes.SetupRoutes(r, myHandlers)

	r.Run(":8080")
}
