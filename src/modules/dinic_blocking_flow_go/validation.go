package mod_dinic_blocking_flow_go

func validatePositiveCap(n int) error {
	if n <= 0 {
		return ErrDinicGoInvalid
	}
	return nil
}
