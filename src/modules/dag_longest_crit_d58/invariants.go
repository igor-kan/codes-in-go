package mod_dag_longest_crit_d58

func verifyPostcondition(ok bool) error {
	if !ok {
		return ErrCritInvariant
	}
	return nil
}
