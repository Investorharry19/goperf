package metrics

import (
	"net/http"
	"time"
)

// Recorder records a completed request.
type Recorder interface {
	Record(method, route string, status int, d time.Duration, at time.Time)
}

// ResponseWriter wraps an http.ResponseWriter and preserves response behavior.
type ResponseWriter struct{ http.ResponseWriter }
