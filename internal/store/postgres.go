package store

import "context"

type postgresStore struct{ dsn string }

func (s *postgresStore) Migrate(ctx context.Context) error { return nil }
func (s *postgresStore) Ping(ctx context.Context) error    { return nil }
func (s *postgresStore) Close() error                      { return nil }
func (s *postgresStore) Driver() string                    { return "postgres" }
func (s *postgresStore) UpsertRequestMetrics(ctx context.Context, rows []RequestRow) error {
	return nil
}
func (s *postgresStore) QueryRequestSeries(ctx context.Context, f RequestFilter, stepSec int64) ([]SeriesPoint, error) {
	return nil, nil
}
func (s *postgresStore) QueryRequestTotals(ctx context.Context, f RequestFilter) (Aggregate, error) {
	return Aggregate{}, nil
}
func (s *postgresStore) QueryRouteStats(ctx context.Context, from, to int64, limit int) ([]RouteStat, error) {
	return nil, nil
}
func (s *postgresStore) InsertRuntimeSample(ctx context.Context, s2 RuntimeSample) error { return nil }
func (s *postgresStore) QueryRuntimeSeries(ctx context.Context, from, to, stepSec int64) ([]RuntimePoint, error) {
	return nil, nil
}
func (s *postgresStore) QueryRuntimeAvg(ctx context.Context, from, to int64) (RuntimePoint, bool, error) {
	return RuntimePoint{}, false, nil
}
func (s *postgresStore) QueryRuntimeLatest(ctx context.Context) (RuntimeSample, bool, error) {
	return RuntimeSample{}, false, nil
}
func (s *postgresStore) InsertLogs(ctx context.Context, rows []LogEntry) error { return nil }
func (s *postgresStore) SearchLogs(ctx context.Context, q LogQuery) ([]LogEntry, error) {
	return nil, nil
}
func (s *postgresStore) CountLogs(ctx context.Context, minLevel int, fromMs, toMs int64) (int64, error) {
	return 0, nil
}
func (s *postgresStore) InsertAlertEvent(ctx context.Context, e AlertEvent) error { return nil }
func (s *postgresStore) ListAlertEvents(ctx context.Context, limit int) ([]AlertEvent, error) {
	return nil, nil
}
func (s *postgresStore) LastAlertEventPerRule(ctx context.Context) (map[string]AlertEvent, error) {
	return map[string]AlertEvent{}, nil
}
func (s *postgresStore) Prune(ctx context.Context, c Cutoffs) (PruneResult, error) {
	return PruneResult{}, nil
}
