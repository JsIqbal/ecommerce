package logger

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type contextKey string

const (
	traceIDKey contextKey = "TraceID"
)

func generateID() string {
	ID, err := uuid.NewRandom()
	if err != nil {
		Error(context.Background(), "cannot generate uuid", err)
		return ""
	}

	return ID.String()
}

func ModifyContext(c *gin.Context) {
	traceID := c.GetHeader("X-Trace-ID")
	if len(traceID) == 0 {
		traceID = generateID()
	}

	data := logData{
		TraceID: traceID,
		SpanID:  generateID(),
	}

	ctx := context.WithValue(c.Request.Context(), traceIDKey, data)
	c.Request = c.Request.WithContext(ctx)
	c.Next()
}

func CreateContext(traceID string) context.Context {
	if len(traceID) == 0 {
		traceID = generateID()
	}

	data := logData{
		TraceID: traceID,
		SpanID:  generateID(),
	}

	ctx := context.WithValue(context.Background(), traceIDKey, data)
	return ctx
}
