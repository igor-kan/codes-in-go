package mod_consistent_bounded_load

func validateRingCfg(servers []string, replicas int, loadFactor float64) error {
	if len(servers) == 0 || replicas <= 0 || loadFactor < 1.0 {
		return ErrBoundedRingInvalid
	}
	return nil
}
