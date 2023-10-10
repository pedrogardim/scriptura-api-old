package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pedrogardim/scriptura-api/api/rest/handlers"
)

func Init(router *gin.Engine) {
	router.GET("/api/verse/:std_ref", handlers.GetVerse)

}
