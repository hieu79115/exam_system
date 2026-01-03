package main

import (
	"exam-system/config"
	"exam-system/internal/app"
	"exam-system/internal/routes"

	"github.com/gin-gonic/gin"
)

// @title           Exam System API
// @version         1.0
// @description     API for managing examinations, questions, and reading passages
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.email  npt911@gmail.com

// @license.name  All rights reserved

// @host      localhost:8080
// @BasePath  /api/v1

// @schemes http https
func main() {
	db := config.InitDB()

	myHandlers := app.InitApp(db)

	r := gin.Default()
	routes.SetupRoutes(r, myHandlers)

	r.Run(":8080")
}
