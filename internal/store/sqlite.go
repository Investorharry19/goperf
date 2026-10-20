package store

import "context"

type sqliteStore struct{ dsn string }

func (s *sqliteStore) Migrate(ctx context.Context) error                                 { return nil }
func (s *sqliteStore) Ping(ctx context.Context) error                                    { return nil }
func (s *sqliteStore) Close() error                                                      { return nil }
func (s *sqliteStore) Driver() string                                                    { return "sqlite" }
func (s *sqliteStore) UpsertRequestMetrics(ctx context.Context, rows []RequestRow) error { return nil }
func (s *sqliteStore) QueryRequestSeries(ctx context.Context, f RequestFilter, stepSec int64) ([]SeriesPoint, error) {
	return nil, nil
}
func (s *sqliteStore) QueryRequestTotals(ctx context.Context, f RequestFilter) (Aggregate, error) {
	return Aggregate{}, nil
}
func (s *sqliteStore) QueryRouteStats(ctx context.Context, from, to int64, limit int) ([]RouteStat, error) {
	return nil, nil
}
func (s *sqliteStore) InsertRuntimeSample(ctx context.Context, s2 RuntimeSample) error { return nil }
func (s *sqliteStore) QueryRuntimeSeries(ctx context.Context, from, to, stepSec int64) ([]RuntimePoint, error) {
	return nil, nil
}
func (s *sqliteStore) QueryRuntimeAvg(ctx context.Context, from, to int64) (RuntimePoint, bool, error) {
	return RuntimePoint{}, false, nil
}
func (s *sqliteStore) QueryRuntimeLatest(ctx context.Context) (RuntimeSample, bool, error) {
	return RuntimeSample{}, false, nil
}
func (s *sqliteStore) InsertLogs(ctx context.Context, rows []LogEntry) error { return nil }
func (s *sqliteStore) SearchLogs(ctx context.Context, q LogQuery) ([]LogEntry, error) {
	return nil, nil
}
func (s *sqliteStore) CountLogs(ctx context.Context, minLevel int, fromMs, toMs int64) (int64, error) {
	return 0, nil
}
func (s *sqliteStore) InsertAlertEvent(ctx context.Context, e AlertEvent) error { return nil }
func (s *sqliteStore) ListAlertEvents(ctx context.Context, limit int) ([]AlertEvent, error) {
	return nil, nil
}
func (s *sqliteStore) LastAlertEventPerRule(ctx context.Context) (map[string]AlertEvent, error) {
	return map[string]AlertEvent{}, nil
}
func (s *sqliteStore) Prune(ctx context.Context, c Cutoffs) (PruneResult, error) {
	return PruneResult{}, nil
}
