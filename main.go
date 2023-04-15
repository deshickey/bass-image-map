package main

import (
	"github.com/deshickey/bass-image-map/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("templates/*.html")

	root := r.Group("/")

	routes.Pages(root)
	r.Run(":443")
}
