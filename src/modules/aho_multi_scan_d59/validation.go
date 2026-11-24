package mod_aho_multi_scan_d59

func validatePositiveCap(n int) error {
	if n <= 0 {
		return ErrAcGoInvalid
	}
	return nil
}
