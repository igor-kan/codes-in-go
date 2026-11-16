package mod_crdt_lww_element_set

import "sort"

func (s *LWWSet) Add(elem string, ts uint64) error {
	if err := ValidateLWWOps(s, elem); err != nil {
		return err
	}
	if prev, ok := s.Adds[elem]; !ok || ts > prev {
		s.Adds[elem] = ts
	}
	return nil
}

func (s *LWWSet) Remove(elem string, ts uint64) error {
	if err := ValidateLWWOps(s, elem); err != nil {
		return err
	}
	if prev, ok := s.Removes[elem]; !ok || ts > prev {
		s.Removes[elem] = ts
	}
	return nil
}

func (s *LWWSet) Contains(elem string) bool {
	addTs, hasAdd := s.Adds[elem]
	if !hasAdd {
		return false
	}
	remTs, hasRem := s.Removes[elem]
	if !hasRem {
		return true
	}
	if addTs > remTs {
		return true
	}
	if remTs > addTs {
		return false
	}
	return s.Bias == BiasAddWins
}

func (s *LWWSet) Merge(other *LWWSet) error {
	if s == nil || other == nil {
		return ErrLWWEmpty
	}
	if s.Bias != other.Bias {
		return ErrLWWInvalid
	}
	for k, ts := range other.Adds {
		if cur, ok := s.Adds[k]; !ok || ts > cur {
			s.Adds[k] = ts
		}
	}
	for k, ts := range other.Removes {
		if cur, ok := s.Removes[k]; !ok || ts > cur {
			s.Removes[k] = ts
		}
	}
	return nil
}

func (s *LWWSet) Elements() []string {
	out := make([]string, 0, len(s.Adds))
	for k := range s.Adds {
		if s.Contains(k) {
			out = append(out, k)
		}
	}
	sort.Strings(out)
	return out
}
