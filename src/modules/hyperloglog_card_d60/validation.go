package mod_hyperloglog_card_d60

func validatePositiveCap(n int) error {
	if n <= 0 {
		return ErrHllEstimatorInvalid
	}
	return nil
}
