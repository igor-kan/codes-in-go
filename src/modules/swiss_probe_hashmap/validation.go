package mod_swiss_probe_hashmap

// ValidateSwissMap checks capacity alignment and slot consistency.
func ValidateSwissMap(m *SwissMap) error {
	if m == nil || m.Capacity == 0 || m.Capacity%GroupWidth != 0 {
		return ErrSwissInvalid
	}
	if len(m.Ctrl) != m.Capacity || len(m.Slots) != m.Capacity {
		return ErrSwissInvariant
	}
	return nil
}
