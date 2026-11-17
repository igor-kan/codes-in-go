package mod_consistent_bounded_load

func verifyBoundedLoad(ring *BoundedHashRing) error {
	loads, tot := ring.LoadSnapshot()
	capLimit := fastMaxAllowed(tot, len(loads), ring.factor) + 1
	sum := 0
	for _, l := range loads {
		if l > capLimit {
			return ErrBoundedRingInvariant
		}
		sum += l
	}
	if sum != tot {
		return ErrBoundedRingInvariant
	}
	return nil
}
