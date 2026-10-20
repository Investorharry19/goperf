package metrics

import "context"

// RequestIDFromContext extracts the request ID from context if present.
func RequestIDFromContext(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	if v, ok := ctx.Value("goperf_request_id").(string); ok {
		return v
	}
	return ""
}
