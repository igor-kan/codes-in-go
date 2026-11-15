package mod_swiss_probe_hashmap

// NewSwissMap creates a power-of-two capacity Swiss-style probing map.
func NewSwissMap(initialGroups int) (*SwissMap, error) {
	if initialGroups < 1 {
		return nil, ErrSwissInvalid
	}
	g := 1
	for g < initialGroups {
		g <<= 1
	}
	capSlots := g * GroupWidth
	ctrl := make([]uint8, capSlots)
	for i := range ctrl {
		ctrl[i] = CtrlEmpty
	}
	return &SwissMap{
		Ctrl:     ctrl,
		Slots:    make([]SwissSlot, capSlots),
		Capacity: capSlots,
	}, nil
}
