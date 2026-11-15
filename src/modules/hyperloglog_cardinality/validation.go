package mod_hyperloglog_cardinality

// ValidateHLL checks register array length against 2^Precision.
func ValidateHLL(h *HLL) error {
	if h == nil || h.Precision < 4 || h.Precision > 16 {
		return ErrHLLInvalid
	}
	if len(h.Registers) != 1<<h.Precision {
		return ErrHLLInvariant
	}
	return nil
}
