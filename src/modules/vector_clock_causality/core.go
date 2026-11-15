package mod_vector_clock_causality

// Tick increments local logical clock for nodeID.
func (vc *VectorClock) Tick(nodeID int) error {
	if nodeID < 0 || nodeID >= len(vc.Ticks) {
		return ErrVClockBounds
	}
	vc.Ticks[nodeID]++
	return nil
}

// Merge takes componentwise max with remote clock and advances receiver nodeID.
func (vc *VectorClock) Merge(remote *VectorClock, receiverID int) error {
	if remote == nil || len(vc.Ticks) != len(remote.Ticks) {
		return ErrVClockInvalid
	}
	if receiverID < 0 || receiverID >= len(vc.Ticks) {
		return ErrVClockBounds
	}
	for i, v := range remote.Ticks {
		if v > vc.Ticks[i] {
			vc.Ticks[i] = v
		}
	}
	vc.Ticks[receiverID]++
	return nil
}

// Compare returns causal partial ordering between vc and other.
func (vc *VectorClock) Compare(other *VectorClock) (CausalRelation, error) {
	if other == nil || len(vc.Ticks) != len(other.Ticks) {
		return RelConcurrent, ErrVClockInvalid
	}
	hasLess, hasGreater := false, false
	for i := range vc.Ticks {
		if vc.Ticks[i] < other.Ticks[i] {
			hasLess = true
		} else if vc.Ticks[i] > other.Ticks[i] {
			hasGreater = true
		}
	}
	switch {
	case !hasLess && !hasGreater:
		return RelEqual, nil
	case hasLess && !hasGreater:
		return RelBefore, nil
	case !hasLess && hasGreater:
		return RelAfter, nil
	default:
		return RelConcurrent, nil
	}
}

// Clone returns an independent copy of vc.
func (vc *VectorClock) Clone() *VectorClock {
	return &VectorClock{Ticks: append([]uint64(nil), vc.Ticks...)}
}
