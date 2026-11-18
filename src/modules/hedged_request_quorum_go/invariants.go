package mod_hedged_request_quorum_go

func verifyPostcondition(ok bool) error {
	if !ok {
		return ErrHedgedQuorumInvariant
	}
	return nil
}
