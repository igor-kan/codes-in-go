package mod_hedged_request_quorum_go

func validatePositiveCap(n int) error {
	if n <= 0 {
		return ErrHedgedQuorumInvalid
	}
	return nil
}
