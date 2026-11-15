package mod_timing_wheel_scheduler

// Schedule registers a timer ID firing after delayTicks >= 1.
func (tw *TimingWheel) Schedule(id, delayTicks uint64, payload string) (*TimerEntry, error) {
	if delayTicks == 0 {
		delayTicks = 1
	}
	if _, dup := tw.Active[id]; dup {
		return nil, ErrTimerInvalid
	}
	ns := uint64(tw.NumSlots)
	rounds := (delayTicks - 1) / ns
	slot := int((tw.CurrentTick + delayTicks) % ns)
	ent := &TimerEntry{
		ID:       id,
		Rounds:   rounds,
		Deadline: tw.CurrentTick + delayTicks,
		Payload:  payload,
		slotIdx:  slot,
		active:   true,
	}
	ent.elem = tw.Buckets[slot].PushBack(ent)
	tw.Active[id] = ent
	return ent, nil
}

// Cancel detaches an active timer by ID.
func (tw *TimingWheel) Cancel(id uint64) bool {
	ent, ok := tw.Active[id]
	if !ok || !ent.active {
		return false
	}
	tw.Buckets[ent.slotIdx].Remove(ent.elem)
	ent.active = false
	delete(tw.Active, id)
	return true
}

// Advance steps the clock forward by steps ticks and returns fired entries in tick order.
func (tw *TimingWheel) Advance(steps uint64) []*TimerEntry {
	var fired []*TimerEntry
	for s := uint64(0); s < steps; s++ {
		tw.CurrentTick++
		slot := int(tw.CurrentTick % uint64(tw.NumSlots))
		b := tw.Buckets[slot]
		for e := b.Front(); e != nil; {
			next := e.Next()
			ent := e.Value.(*TimerEntry)
			if ent.Rounds > 0 {
				ent.Rounds--
			} else {
				b.Remove(e)
				ent.active = false
				delete(tw.Active, ent.ID)
				fired = append(fired, ent)
			}
			e = next
		}
	}
	return fired
}
