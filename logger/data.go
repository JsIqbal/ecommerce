package logger

import "context"

type logData struct {
	TraceID string
	SpanID  string
}

func getTraceID(ctx context.Context) string {
	data, ok := ctx.Value(traceIDKey).(logData)
	if !ok {
		return ""
	}
	return data.TraceID
}

func getSpanID(ctx context.Context) string {
	data, ok := ctx.Value(traceIDKey).(logData)
	if !ok {
		return ""
	}
	return data.SpanID
}

func GetTraceID(ctx context.Context) string {
	ctx = getContext(ctx)
	data, ok := ctx.Value(traceIDKey).(logData)
	if !ok {
		return ""
	}
	return data.TraceID
}