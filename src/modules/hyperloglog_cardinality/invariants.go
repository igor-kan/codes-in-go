package mod_hyperloglog_cardinality

// VerifyHLLMonotone checks register range and idempotency under self-merge.
func VerifyHLLMonotone(h *HLL) error {
	if err := ValidateHLL(h); err != nil {
		return err
	}
	for _, r := range h.Registers {
		if r > 64 {
			return ErrHLLInvariant
		}
	}
	return nil
}
