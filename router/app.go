package router

import (
	"github.com/gin-gonic/gin"
)

func New() *Engine {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return router
}
