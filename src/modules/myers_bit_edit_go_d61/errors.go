package mod_myers_bit_edit_go_d61

import "errors"

var (
	ErrMyersEditInvalid   = errors.New("myers_bit_edit_go_d61: invalid argument")
	ErrMyersEditBounds    = errors.New("myers_bit_edit_go_d61: index out of bounds")
	ErrMyersEditEmpty     = errors.New("myers_bit_edit_go_d61: empty input")
	ErrMyersEditInvariant = errors.New("myers_bit_edit_go_d61: invariant violated")
)
