package mod_union_find_weighted

// ComponentCount counts distinct root representatives.
func (d *WeightedDSU) ComponentCount() int {
	c := 0
	for i, p := range d.Parent {
		if i == p {
			c++
		}
	}
	return c
}
