package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic recovered: %v", err)
				c.JSON(500, gin.H{
					"code":    500,
					"message": "Internal server error",
					"data":    nil,
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}
