package mod_tarjan_scc_condense_go

import "errors"

var (
	ErrTarjanGoInvalid   = errors.New("tarjan_scc_condense_go: invalid argument")
	ErrTarjanGoBounds    = errors.New("tarjan_scc_condense_go: index out of bounds")
	ErrTarjanGoEmpty     = errors.New("tarjan_scc_condense_go: empty input")
	ErrTarjanGoInvariant = errors.New("tarjan_scc_condense_go: invariant violated")
)
