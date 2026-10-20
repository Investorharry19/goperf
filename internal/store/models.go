package store

import "context"

// RequestRow stores one aggregated request metric row.
type RequestRow struct {
	BucketStart int64
	Method      string
	Route       string
	Status      int
	Count       int64
	SumMs       float64
	MaxMs       float64
	Hist        [14]int64
}

// RequestFilter filters request metric queries.
type RequestFilter struct {
	From, To int64
	Route    string
	Method   string
}

// Aggregate holds aggregated request metric data.
type Aggregate struct {
	Requests, Errors5xx, Errors4xx int64
	SumMs, MaxMs                   float64
	Hist                           [14]int64
}

// SeriesPoint is one step in a time series.
type SeriesPoint struct {
	T int64
	Aggregate
}

// RouteStat is a per-route summary.
type RouteStat struct {
	Method string
	Route  string
	Aggregate
}

// RuntimeSample stores one runtime snapshot.
type RuntimeSample struct {
	Instance   string
	TS         int64
	HeapBytes  int64
	SysBytes   int64
	Goroutines int64
}

// RuntimePoint stores averaged runtime values.
type RuntimePoint struct {
	T                               int64
	HeapBytes, SysBytes, Goroutines float64
}

// LogEntry is one stored log entry.
type LogEntry struct {
	ID        int64
	TSMs      int64
	Level     int
	LevelName string
	Message   string
	Attrs     string
	RequestID string
}

// LogQuery filters log searches.
type LogQuery struct {
	MinLevel  *int
	Q         string
	RequestID string
	FromMs    int64
	ToMs      int64
	BeforeID  int64
	Limit     int
}

// AlertEvent is one stored alert state change.
type AlertEvent struct {
	ID        int64
	TSMs      int64
	Rule      string
	State     string
	Value     float64
	Threshold float64
	Message   string
}

// Cutoffs holds retention cutoffs.
type Cutoffs struct {
	MetricsBefore  int64
	RuntimeBefore  int64
	LogsBeforeMs   int64
	AlertsBeforeMs int64
}

// PruneResult reports deleted rows during pruning.
type PruneResult struct {
	Metrics int64
	Runtime int64
	Logs    int64
	Alerts  int64
}

// Store provides persistence for goperf data.
type Store interface {
	Migrate(ctx context.Context) error
	Ping(ctx context.Context) error
	Close() error
	Driver() string
	UpsertRequestMetrics(ctx context.Context, rows []RequestRow) error
	QueryRequestSeries(ctx context.Context, f RequestFilter, stepSec int64) ([]SeriesPoint, error)
	QueryRequestTotals(ctx context.Context, f RequestFilter) (Aggregate, error)
	QueryRouteStats(ctx context.Context, from, to int64, limit int) ([]RouteStat, error)
	InsertRuntimeSample(ctx context.Context, s RuntimeSample) error
	QueryRuntimeSeries(ctx context.Context, from, to, stepSec int64) ([]RuntimePoint, error)
	QueryRuntimeAvg(ctx context.Context, from, to int64) (RuntimePoint, bool, error)
	QueryRuntimeLatest(ctx context.Context) (RuntimeSample, bool, error)
	InsertLogs(ctx context.Context, rows []LogEntry) error
	SearchLogs(ctx context.Context, q LogQuery) ([]LogEntry, error)
	CountLogs(ctx context.Context, minLevel int, fromMs, toMs int64) (int64, error)
	InsertAlertEvent(ctx context.Context, e AlertEvent) error
	ListAlertEvents(ctx context.Context, limit int) ([]AlertEvent, error)
	LastAlertEventPerRule(ctx context.Context) (map[string]AlertEvent, error)
	Prune(ctx context.Context, c Cutoffs) (PruneResult, error)
}
