package mod_deadline_priority_heap

type TimerTask struct {
	ID       string
	Deadline int64
	index    int
}
