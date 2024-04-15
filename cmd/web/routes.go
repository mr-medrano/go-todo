package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func routes() http.Handler {
	e := gin.New()
	e.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	return e
}
