package routes

import (
	"github.com/deshickey/bass-image-map/handlers"
	"github.com/gin-gonic/gin"
)

func Pages(g *gin.RouterGroup) {
	g.GET("/", handlers.GetIndex())
	g.POST("/scale/:note/:scale", handlers.GetScale())
}
