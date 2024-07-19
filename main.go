package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"git.kasikornline.com/pdf-decrypt/internal/middleware"
	"git.kasikornline.com/pdf-decrypt/internal/router"
)

func main() {

	_ = godotenv.Load(".env")
	port := os.Getenv("PORT")

	//initz gin
	engine := gin.Default()
	engine.Use(middleware.CORSMiddleware())
	router.Setup(engine)

	_ = engine.Run(":" + port)

}
