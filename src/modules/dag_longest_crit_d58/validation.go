package mod_dag_longest_crit_d58

func validatePositiveCap(n int) error {
	if n <= 0 {
		return ErrCritInvalid
	}
	return nil
}
