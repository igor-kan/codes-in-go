package mod_interval_coalesce_go

type IntervalLedger interface {
	Add(start, end int64) error
	Contains(x int64) bool
	Snapshots() []Span
}
