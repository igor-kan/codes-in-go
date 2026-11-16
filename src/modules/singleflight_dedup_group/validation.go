package mod_singleflight_dedup_group

func ValidateFlightKey(key string) error {
	if key == "" {
		return ErrFlightEmpty
	}
	return nil
}
