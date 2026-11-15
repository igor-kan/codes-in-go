package mod_crdt_gcounter_pncounter

// EqualPNCounters checks structural equality of replica vectors.
func EqualPNCounters(a, b *PNCounter) bool {
	if len(a.P.Slots) != len(b.P.Slots) || len(a.N.Slots) != len(b.N.Slots) {
		return false
	}
	for i := range a.P.Slots {
		if a.P.Slots[i] != b.P.Slots[i] || a.N.Slots[i] != b.N.Slots[i] {
			return false
		}
	}
	return true
}
