package rest

import (
	"github.com/gin-gonic/gin"
)

func corsMiddleware(c *gin.Context) {
	// Allow all origins
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	// Allow specific headers
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")

	// Allow all methods
	c.Writer.Header().Set("Access-Control-Allow-Methods", "*")

	// Handle OPTIONS method
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}

	c.Next()
}
