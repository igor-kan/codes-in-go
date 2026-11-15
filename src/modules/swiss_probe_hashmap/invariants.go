package mod_swiss_probe_hashmap

// VerifySwissCounts ensures live slot and tombstone metadata counts are exact.
func VerifySwissCounts(m *SwissMap) error {
	if err := ValidateSwissMap(m); err != nil {
		return err
	}
	live, tomb := 0, 0
	for _, c := range m.Ctrl {
		switch {
		case c == CtrlDeleted:
			tomb++
		case c&0x80 == 0:
			live++
		}
	}
	if live != m.Count || tomb != m.Tomb {
		return ErrSwissInvariant
	}
	return nil
}
