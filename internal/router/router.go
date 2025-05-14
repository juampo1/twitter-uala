package router

import (
	"net/http"
	"twitter-uala/internal/models"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusText(http.StatusOK),
		})
	})

	//create a tweet
	router.POST("{:userId}/tweet", func(c *gin.Context) {
		var tweet models.Tweet
		if err := c.ShouldBindJSON(&tweet); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	})

	return router
}
