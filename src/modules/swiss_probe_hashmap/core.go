package mod_swiss_probe_hashmap

func swissHash(key uint64) uint64 {
	key = (key ^ (key >> 30)) * 0xbf58476d1ce4e5b9
	key = (key ^ (key >> 27)) * 0x94d049bb133111eb
	return key ^ (key >> 31)
}

func splitH1H2(h uint64) (uint64, uint8) {
	return h >> 7, uint8(h & 0x7F)
}

func (m *SwissMap) rehash(newCap int) {
	oldCtrl := m.Ctrl
	oldSlots := m.Slots
	m.Ctrl = make([]uint8, newCap)
	for i := range m.Ctrl {
		m.Ctrl[i] = CtrlEmpty
	}
	m.Slots = make([]SwissSlot, newCap)
	m.Capacity = newCap
	m.Count = 0
	m.Tomb = 0
	for i, c := range oldCtrl {
		if c&0x80 == 0 {
			m.Put(oldSlots[i].Key, oldSlots[i].Val)
		}
	}
}

// Put inserts or overwrites key -> val with automatic 7/8 load resize.
func (m *SwissMap) Put(key, val uint64) {
	if (m.Count+m.Tomb+1)*8 >= m.Capacity*7 {
		newCap := m.Capacity
		if m.Count*2 >= m.Capacity {
			newCap *= 2
		}
		m.rehash(newCap)
	}
	h := swissHash(key)
	h1, h2 := splitH1H2(h)
	mask := uint64(m.Capacity/GroupWidth - 1)
	g := h1 & mask
	firstTomb := -1
	for step := uint64(0); ; step++ {
		base := int(g) * GroupWidth
		for j := 0; j < GroupWidth; j++ {
			pos := base + j
			c := m.Ctrl[pos]
			if c == h2 && m.Slots[pos].Key == key {
				m.Slots[pos].Val = val
				return
			}
			if c == CtrlDeleted && firstTomb < 0 {
				firstTomb = pos
			}
			if c == CtrlEmpty {
				ins := pos
				if firstTomb >= 0 {
					ins = firstTomb
					m.Tomb--
				}
				m.Ctrl[ins] = h2
				m.Slots[ins] = SwissSlot{Key: key, Val: val}
				m.Count++
				return
			}
		}
		g = (g + step + 1) & mask
	}
}

// Get retrieves val for key if present.
func (m *SwissMap) Get(key uint64) (uint64, bool) {
	h1, h2 := splitH1H2(swissHash(key))
	mask := uint64(m.Capacity/GroupWidth - 1)
	g := h1 & mask
	for step := uint64(0); step <= mask; step++ {
		base := int(g) * GroupWidth
		for j := 0; j < GroupWidth; j++ {
			pos := base + j
			c := m.Ctrl[pos]
			if c == h2 && m.Slots[pos].Key == key {
				return m.Slots[pos].Val, true
			}
			if c == CtrlEmpty {
				return 0, false
			}
		}
		g = (g + step + 1) & mask
	}
	return 0, false
}

// Delete tombstones key if present and returns true if deleted.
func (m *SwissMap) Delete(key uint64) bool {
	h1, h2 := splitH1H2(swissHash(key))
	mask := uint64(m.Capacity/GroupWidth - 1)
	g := h1 & mask
	for step := uint64(0); step <= mask; step++ {
		base := int(g) * GroupWidth
		for j := 0; j < GroupWidth; j++ {
			pos := base + j
			c := m.Ctrl[pos]
			if c == h2 && m.Slots[pos].Key == key {
				m.Ctrl[pos] = CtrlDeleted
				m.Count--
				m.Tomb++
				return true
			}
			if c == CtrlEmpty {
				return false
			}
		}
		g = (g + step + 1) & mask
	}
	return false
}
