package logger

import (
	"context"

	"github.com/gin-gonic/gin"
)

func getContext(ctx interface{}) context.Context {
	switch ctx := ctx.(type) {
	case *gin.Context:
		return ctx.Request.Context()
	case context.Context:
		return ctx
	default:
		return context.Background()
	}
}
