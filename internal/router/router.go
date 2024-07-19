package router

import (
	"github.com/gin-gonic/gin"

	"git.kasikornline.com/pdf-decrypt/internal/controller"
)

func Setup(router *gin.Engine) {
	router.POST("/decrypt-pdf", controller.DecryptPdfFile)
	router.GET("/book", controller.GetBook)
}
