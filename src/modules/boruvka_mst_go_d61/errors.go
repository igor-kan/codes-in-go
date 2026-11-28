package mod_boruvka_mst_go_d61

import "errors"

var (
	ErrBoruvkaMstInvalid   = errors.New("boruvka_mst_go_d61: invalid argument")
	ErrBoruvkaMstBounds    = errors.New("boruvka_mst_go_d61: index out of bounds")
	ErrBoruvkaMstEmpty     = errors.New("boruvka_mst_go_d61: empty input")
	ErrBoruvkaMstInvariant = errors.New("boruvka_mst_go_d61: invariant violated")
)
