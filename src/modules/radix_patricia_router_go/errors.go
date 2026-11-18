package mod_radix_patricia_router_go

import "errors"

var (
	ErrRadixRouterInvalid   = errors.New("radix_patricia_router_go: invalid argument")
	ErrRadixRouterBounds    = errors.New("radix_patricia_router_go: index out of bounds")
	ErrRadixRouterEmpty     = errors.New("radix_patricia_router_go: empty input")
	ErrRadixRouterInvariant = errors.New("radix_patricia_router_go: invariant violated")
)
