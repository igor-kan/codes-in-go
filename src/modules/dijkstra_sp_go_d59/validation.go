package mod_dijkstra_sp_go_d59

func validatePositiveCap(n int) error {
	if n <= 0 {
		return ErrSsspInvalid
	}
	return nil
}
