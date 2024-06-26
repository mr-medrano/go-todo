package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"

	"examples.mrmedano.todo/internal/schemas"
)

func (a *Application) ping(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "pong"})
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

	c.IndentedJSON(http.StatusCreated, gin.H{"id": id})
}

func (a *Application) taskView(c *gin.Context) {
	id := c.Param("id")

	task, err := a.tasks.Get(c, id)
	if err != nil {
		if err == pgx.ErrNoRows {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "INTERNALERR-1",
			"message": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, task)
}
