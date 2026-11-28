package mod_douglas_peucker_go_d61

func validatePositiveCap(n int) error {
	if n <= 0 {
		return ErrPolylineRdpInvalid
	}
	return nil
}
