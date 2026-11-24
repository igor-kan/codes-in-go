package mod_kmp_stream_go_d59

func validatePositiveCap(n int) error {
	if n <= 0 {
		return ErrKmpGoInvalid
	}
	return nil
}
