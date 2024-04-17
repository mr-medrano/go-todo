package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *Application) ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}
