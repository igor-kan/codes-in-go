package mod_crdt_gcounter_pncounter

// VerifySemilatticeLaws checks commutativity, associativity, and idempotency on (a, b, c).
func VerifySemilatticeLaws(a, b, c *PNCounter) error {
	ab := a.Clone()
	_ = ab.Merge(b)
	ba := b.Clone()
	_ = ba.Merge(a)
	if !EqualPNCounters(ab, ba) {
		return ErrCRDTInvariant
	}
	aa := a.Clone()
	_ = aa.Merge(a)
	if !EqualPNCounters(a, aa) {
		return ErrCRDTInvariant
	}
	abc1 := ab.Clone()
	_ = abc1.Merge(c)
	bc := b.Clone()
	_ = bc.Merge(c)
	abc2 := a.Clone()
	_ = abc2.Merge(bc)
	if !EqualPNCounters(abc1, abc2) {
		return ErrCRDTInvariant
	}
	return nil
}
