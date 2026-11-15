package mod_crdt_gcounter_pncounter

// Inc adds delta >= 0 to replica slot rid.
func (g *GCounter) Inc(rid int, delta uint64) error {
	if rid < 0 || rid >= len(g.Slots) {
		return ErrCRDTBounds
	}
	g.Slots[rid] += delta
	return nil
}

// Value returns the global sum of all replica slots.
func (g *GCounter) Value() uint64 {
	var sum uint64
	for _, v := range g.Slots {
		sum += v
	}
	return sum
}

// Join merges other into g by taking componentwise maximum.
func (g *GCounter) Join(other *GCounter) error {
	if other == nil || len(g.Slots) != len(other.Slots) {
		return ErrCRDTInvalid
	}
	for i, v := range other.Slots {
		if v > g.Slots[i] {
			g.Slots[i] = v
		}
	}
	return nil
}

// Add adjusts replica rid by signed delta (negative routes to N counter).
func (pn *PNCounter) Add(rid int, delta int64) error {
	if delta >= 0 {
		return pn.P.Inc(rid, uint64(delta))
	}
	return pn.N.Inc(rid, uint64(-delta))
}

// Value returns P.Value() - N.Value().
func (pn *PNCounter) Value() int64 {
	return int64(pn.P.Value()) - int64(pn.N.Value())
}

// Merge joins both P and N GCounters.
func (pn *PNCounter) Merge(other *PNCounter) error {
	if other == nil {
		return ErrCRDTEmpty
	}
	if err := pn.P.Join(&other.P); err != nil {
		return err
	}
	return pn.N.Join(&other.N)
}

// Clone deep-copies pn.
func (pn *PNCounter) Clone() *PNCounter {
	cp := &PNCounter{
		P: GCounter{Slots: append([]uint64(nil), pn.P.Slots...)},
		N: GCounter{Slots: append([]uint64(nil), pn.N.Slots...)},
	}
	return cp
}
