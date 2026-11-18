package mod_dinic_blocking_flow_go

import "errors"

var (
	ErrDinicGoInvalid   = errors.New("dinic_blocking_flow_go: invalid argument")
	ErrDinicGoBounds    = errors.New("dinic_blocking_flow_go: index out of bounds")
	ErrDinicGoEmpty     = errors.New("dinic_blocking_flow_go: empty input")
	ErrDinicGoInvariant = errors.New("dinic_blocking_flow_go: invariant violated")
)
