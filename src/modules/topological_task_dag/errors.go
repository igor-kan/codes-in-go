package mod_topological_task_dag

import "errors"

var (
	ErrDAGEmpty     = errors.New("dag: empty or nil input")
	ErrDAGBounds    = errors.New("dag: index or dimension out of bounds")
	ErrDAGInvalid   = errors.New("dag: invalid configuration or state")
	ErrDAGInvariant = errors.New("dag: structural invariant violated")
)
