package mod_crdt_lww_element_set

func (s *LWWSet) Clone() *LWWSet {
	cp := &LWWSet{
		Bias:    s.Bias,
		Adds:    make(map[string]uint64, len(s.Adds)),
		Removes: make(map[string]uint64, len(s.Removes)),
	}
	for k, v := range s.Adds {
		cp.Adds[k] = v
	}
	for k, v := range s.Removes {
		cp.Removes[k] = v
	}
	return cp
}
