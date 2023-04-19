package contexts

import (
	"context"
)

func SetInt64(ctx context.Context, key ContextType, val int64) context.Context {
	return context.WithValue(ctx, key, val)
}

func Int64(ctx context.Context, key ContextType) int64 {
	if ctx == nil {
		return 0
	}
	valInterface := ctx.Value(key)
	if val, ok := valInterface.(int64); ok {
		return val
	}
	return 0
}
