package mod_gale_shapley_stable_go_d61

import "errors"

var (
	ErrStableMarriageInvalid   = errors.New("gale_shapley_stable_go_d61: invalid argument")
	ErrStableMarriageBounds    = errors.New("gale_shapley_stable_go_d61: index out of bounds")
	ErrStableMarriageEmpty     = errors.New("gale_shapley_stable_go_d61: empty input")
	ErrStableMarriageInvariant = errors.New("gale_shapley_stable_go_d61: invariant violated")
)
