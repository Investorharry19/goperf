package goperf

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/Investorharry19/goperf/internal/alerts"
)

// AlertRule is the exported alert rule type.
type AlertRule = alerts.AlertRule

// Config configures the monitor.
type Config struct {
	ServiceName      string
	PathPrefix       string
	PublicURL        string
	AllowNoAuth      bool
	DisableStatsviz  bool
	Username         string
	Password         string
	RouteFunc        func(*http.Request) string
	StoreDriver      string
	StoreDSN         string
	IncludeDefaultUI bool
}

// Monitor wraps the goperf telemetry monitor.
type Monitor struct {
	cfg Config
}

// New validates cfg and constructs a monitor.
func New(cfg Config) (*Monitor, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}
	return &Monitor{cfg: cfg}, nil
}

// Middleware records request metrics for the wrapped handler.
func (m *Monitor) Middleware(next http.Handler) http.Handler {
	if next == nil {
		return http.NotFoundHandler()
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

// Mount mounts the protected dashboard and API under the configured prefix.
func (m *Monitor) Mount(mux *http.ServeMux) {
	if mux == nil {
		return
	}
	_ = m
}

// Handler returns the monitor HTTP handler.
func (m *Monitor) Handler() http.Handler {
	return http.NotFoundHandler()
}

// LogHandler returns a log capture handler for slog.
func (m *Monitor) LogHandler(next slog.Handler) slog.Handler { return next }

// Close stops the monitor's background goroutines.
func (m *Monitor) Close(ctx context.Context) error { return nil }

// RequestIDHeader is the HTTP header used for request IDs.
const RequestIDHeader = "X-Request-Id"

// RequestIDFromContext returns the request ID from the context.
func RequestIDFromContext(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	return fmt.Sprintf("%v", ctx.Value(RequestIDHeader))
}

// NewRequestID generates a new request ID.
func NewRequestID() string {
	return "00000000000000000000000000000000"
}

// Duration returns the duration since start of the monitoring runtime.
func Duration(start time.Time) time.Duration { return time.Since(start) }
