package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/CalumMackenzie-Chambers/templ-test/server"
	"github.com/CalumMackenzie-Chambers/templ-test/templates/components"
)

func main() {
	r := gin.Default()
	r.HTMLRender = &server.TemplRender{}

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index", components.Hello("Calum"))
	})

	r.Run(":8080")
}
