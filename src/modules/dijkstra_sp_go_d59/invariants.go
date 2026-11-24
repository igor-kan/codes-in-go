package mod_dijkstra_sp_go_d59

func verifyPostcondition(ok bool) error {
	if !ok {
		return ErrSsspInvariant
	}
	return nil
}
