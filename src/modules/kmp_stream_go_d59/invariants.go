package mod_kmp_stream_go_d59

func verifyPostcondition(ok bool) error {
	if !ok {
		return ErrKmpGoInvariant
	}
	return nil
}
