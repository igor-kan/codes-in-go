package mod_hedged_request_quorum_go

import "errors"

var (
	ErrHedgedQuorumInvalid   = errors.New("hedged_request_quorum_go: invalid argument")
	ErrHedgedQuorumBounds    = errors.New("hedged_request_quorum_go: index out of bounds")
	ErrHedgedQuorumEmpty     = errors.New("hedged_request_quorum_go: empty input")
	ErrHedgedQuorumInvariant = errors.New("hedged_request_quorum_go: invariant violated")
)
