package contexts

import (
	"context"
)

func SetString(ctx context.Context, key ContextType, val string) context.Context {
	return context.WithValue(ctx, key, val)
}

func String(ctx context.Context, key ContextType) string {
	if ctx == nil {
		return ""
	}
	valInterface := ctx.Value(key)
	if val, ok := valInterface.(string); ok {
		return val
	}
	return ""
}
