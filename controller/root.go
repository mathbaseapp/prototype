package controller

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"prototype.mathbase.app/service"
)

// RunServer サーバーを立ち上げる
func RunServer() {

	engine := gin.Default()

	// cors対応
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"https://mathbase.app"}
	engine.Use(cors.New(config))

	engine.LoadHTMLGlob("template/*.html")
	renderer := multitemplate.NewRenderer()
	engine.GET("/", index)
	renderer.AddFromFiles("index", "template/layouts.html", "template/index.html")
	engine.GET("/search", queryByTex)
	renderer.AddFromFiles("search", "template/layouts.html", "template/search.html")
	engine.HTMLRender = renderer
	engine.GET("/api/v1/search", restQueryByTex)

	engine.Run(":3000")
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index", gin.H{})
}

func queryByTex(c *gin.Context) {
	queryStr := c.DefaultQuery("query", "")
	results, err := service.QueryByLatex(queryStr)
	if err != nil {
		panic(err)
	}
	c.HTML(http.StatusOK, "search", gin.H{
		"query":   queryStr,
		"results": results,
	})
}

func restQueryByTex(c *gin.Context) {
	queryStr := c.DefaultQuery("query", "")
	results, err := service.QueryByLatex(queryStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"results": results,
		"query":   queryStr,
	})
}
