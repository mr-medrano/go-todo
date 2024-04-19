package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"examples.mrmedano.todo/internal/schemas"
)

func (a *Application) ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func (a *Application) taskCreate(c *gin.Context) {
	var task schemas.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "VALIDATEERR-1",
			"message": "Invalid inputs",
		})
		return
	}

	id, err := a.tasks.Insert(c, task.Title, task.Note)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "INTERNALERR-1",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (a *Application) taskView(c *gin.Context) {
	id := c.Param("id")

	task, err := a.tasks.Get(c, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "INTERNALERR-1",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, task)
}
