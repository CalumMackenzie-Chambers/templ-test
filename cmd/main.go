package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/CalumMackenzie-Chambers/templ-test/internal/websocket"
	"github.com/CalumMackenzie-Chambers/templ-test/server"
	"github.com/CalumMackenzie-Chambers/templ-test/server/middleware"
	"github.com/CalumMackenzie-Chambers/templ-test/templates/layouts"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	goenv := os.Getenv("GOENV")
	if goenv == "development" {
		r.Use(middleware.NoCacheMiddleware()) // Apply the middleware only in development mode
	}

	r.HTMLRender = &server.TemplRender{}
	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index", layouts.BaseLayout())
	})

	if goenv == "development" {
		r.GET("/ws", websocket.WSHandler)
		r.POST("/trigger", websocket.TriggerHandler)

		go func() {
			maxRetries := 20
			retries := 0
			for retries < maxRetries {
				if websocket.IsWsConnected() {
					websocket.TriggerRefresh()
					break
				}
				retries++
				time.Sleep(100 * time.Millisecond)
			}
			if retries >= maxRetries {
				log.Println("Max retries reached, could not trigger refresh")
			}
		}()
	}

	r.Run(":8080")
}
