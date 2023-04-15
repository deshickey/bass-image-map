package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/deshickey/bass-image-map/music"
	"github.com/gin-gonic/gin"
)

func GetIndex() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"content": "Handling index page.",
		})
	}
}

func GetScale() gin.HandlerFunc {
	return func(c *gin.Context) {
		root := c.Param("note")
		scale := c.Param("scale")

		scaleIndex, err := strconv.Atoi(scale)
		if err != nil {
			fmt.Printf("failed to convert string to int: %v", err)
		}

		payload := music.GetScale(scaleIndex, root)
		c.JSON(http.StatusOK, payload)
	}
}
