package controller

import (
	"net/http"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

// RunServer サーバーを立ち上げる
func RunServer() {

	engine := gin.Default()
	engine.LoadHTMLGlob("template/*.html")
	renderer := multitemplate.NewRenderer()
	engine.GET("/", index)
	renderer.AddFromFiles("index", "template/layouts.html", "template/index.html")
	engine.GET("/search", queryByTex)
	renderer.AddFromFiles("search", "template/layouts.html", "template/search.html")
	engine.HTMLRender = renderer
	engine.Run(":3000")
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index", gin.H{})
}

func queryByTex(c *gin.Context) {
	results := []PageResult{
		PageResult{"qiita-home1", "https://qiita.com"},
		PageResult{"qiita-home2", "https://qiita.com"},
	}
	queryStr := c.DefaultQuery("query", "")
	c.HTML(http.StatusOK, "search", gin.H{
		"query":   queryStr,
		"results": results,
	})
}
