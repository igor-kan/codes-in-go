package mod_booth_min_rotation_go_d61

import "errors"

var (
	ErrBoothRotationInvalid   = errors.New("booth_min_rotation_go_d61: invalid argument")
	ErrBoothRotationBounds    = errors.New("booth_min_rotation_go_d61: index out of bounds")
	ErrBoothRotationEmpty     = errors.New("booth_min_rotation_go_d61: empty input")
	ErrBoothRotationInvariant = errors.New("booth_min_rotation_go_d61: invariant violated")
)
