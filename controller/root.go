package controller

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"prototype.mathbase.app/lg"
	"prototype.mathbase.app/service"
)

// RunServer サーバーを立ち上げる
func RunServer() {

	engine := gin.Default()

	// cors対応
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"https://mathbase.app", "http://localhost:8080"}
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
	results, queries, err := service.QueryByLatex(queryStr)
	if err != nil {
		lg.I.Println(err)
		c.HTML(http.StatusBadRequest, "index", gin.H{
			"error": "something wrong with processing the query.: \"" + queryStr + "\"",
		})
	} else {
		c.HTML(http.StatusOK, "search", gin.H{
			"results": results,
			"queries": queries,
			"input":   queryStr,
		})
	}
}

func restQueryByTex(c *gin.Context) {
	queryStr := c.DefaultQuery("query", "")
	results, _, err := service.QueryByLatex(queryStr)
	if err != nil {
		lg.I.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "something wrong with processing the query.: \"" + queryStr + "\"",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"results": results,
			"query":   queryStr,
		})
	}
}
