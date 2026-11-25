package mod_hyperloglog_card_d60

func verifyPostcondition(ok bool) error {
	if !ok {
		return ErrHllEstimatorInvariant
	}
	return nil
}
