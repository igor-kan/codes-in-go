package mod_crdt_gcounter_pncounter

// ValidatePNCounter ensures P and N replica dimensions match and are non-empty.
func ValidatePNCounter(pn *PNCounter) error {
	if pn == nil || len(pn.P.Slots) == 0 || len(pn.P.Slots) != len(pn.N.Slots) {
		return ErrCRDTInvalid
	}
	return nil
}
