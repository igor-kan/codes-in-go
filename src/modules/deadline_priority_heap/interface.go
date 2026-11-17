package mod_deadline_priority_heap

type TimerScheduler interface {
	Schedule(id string, deadline int64) error
	Cancel(id string) bool
	DrainExpired(now int64) []string
}
