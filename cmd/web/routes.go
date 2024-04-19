package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *Application) routes() http.Handler {
	e := gin.Default()
	e.GET("/ping", a.ping)
	e.GET("/tasks/:id", a.taskView)
	e.POST("/tasks", a.taskCreate)

	return e
}
