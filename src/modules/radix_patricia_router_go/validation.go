package mod_radix_patricia_router_go

func validatePositiveCap(n int) error {
	if n <= 0 {
		return ErrRadixRouterInvalid
	}
	return nil
}
