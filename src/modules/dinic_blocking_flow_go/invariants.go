package mod_dinic_blocking_flow_go

func verifyPostcondition(ok bool) error {
	if !ok {
		return ErrDinicGoInvariant
	}
	return nil
}
