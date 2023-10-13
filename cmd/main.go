package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/CalumMackenzie-Chambers/templ-test/server"
	"github.com/CalumMackenzie-Chambers/templ-test/templates/layouts"
)

func main() {
	r := gin.Default()
	r.HTMLRender = &server.TemplRender{}
	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index", layouts.BaseLayout())
	})

	r.Run(":8080")
}
