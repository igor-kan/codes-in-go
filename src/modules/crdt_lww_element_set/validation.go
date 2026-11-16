package mod_crdt_lww_element_set

func ValidateLWWOps(s *LWWSet, elem string) error {
	if s == nil || elem == "" {
		return ErrLWWEmpty
	}
	return nil
}
