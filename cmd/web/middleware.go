package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func requestIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := uuid.New().String()
		c.Header("X-Request-ID", uuid)
		c.Next()
	}
}
