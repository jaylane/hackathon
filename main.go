package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"jay-lane.com/hackathon/gpt"
	"jay-lane.com/hackathon/pdf"
	"jay-lane.com/hackathon/structs"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/create/downloadable", func(c *gin.Context) {
		var downloadable structs.Downloadable

		err := c.BindJSON(&downloadable)
		err = gpt.SendDownloadableToChatGPT(downloadable)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		} else {
			pdf.BuildPDF(downloadable)
			c.JSON(http.StatusOK, gin.H{
				"message": "downloadable created",
			})
		}
	})
	r.GET("/test-page", func(c *gin.Context) {
		pdf.TestPage()
		c.JSON(http.StatusOK, gin.H{
			"message": "test pdf generated",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
