package mod_union_find_weighted

// VerifyNoSelfLoops asserts root weights stay identically 0 after path compression.
func VerifyNoSelfLoops(d *WeightedDSU) error {
	for i := range d.Parent {
		_ = d.Find(i)
	}
	return ValidateWeightedDSU(d)
}
