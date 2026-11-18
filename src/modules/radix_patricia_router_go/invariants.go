package mod_radix_patricia_router_go

func verifyPostcondition(ok bool) error {
	if !ok {
		return ErrRadixRouterInvariant
	}
	return nil
}
