package mod_union_find_weighted

// ValidateWeightedDSU verifies array lengths and root self-weights.
func ValidateWeightedDSU(d *WeightedDSU) error {
	if d == nil || len(d.Parent) == 0 || len(d.Weight) != len(d.Parent) {
		return ErrWDSUInvalid
	}
	for i, p := range d.Parent {
		if p < 0 || p >= len(d.Parent) || (p == i && d.Weight[i] != 0) {
			return ErrWDSUInvariant
		}
	}
	return nil
}
