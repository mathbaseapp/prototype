package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RunServer サーバーを立ち上げる
func RunServer() {

	engine := gin.Default()
	engine.GET("/search", queryByTex)
	engine.Run(":3000")
}

func queryByTex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hey!!",
	})
}
