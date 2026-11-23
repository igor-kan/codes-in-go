package mod_dag_longest_crit_d58

import "errors"

var (
	ErrCritInvalid   = errors.New("dag_longest_crit_d58: invalid argument")
	ErrCritBounds    = errors.New("dag_longest_crit_d58: index out of bounds")
	ErrCritEmpty     = errors.New("dag_longest_crit_d58: empty input")
	ErrCritInvariant = errors.New("dag_longest_crit_d58: invariant violated")
)
