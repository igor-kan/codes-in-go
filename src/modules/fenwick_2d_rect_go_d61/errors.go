package mod_fenwick_2d_rect_go_d61

import "errors"

var (
	ErrFenwick2DInvalid   = errors.New("fenwick_2d_rect_go_d61: invalid argument")
	ErrFenwick2DBounds    = errors.New("fenwick_2d_rect_go_d61: index out of bounds")
	ErrFenwick2DEmpty     = errors.New("fenwick_2d_rect_go_d61: empty input")
	ErrFenwick2DInvariant = errors.New("fenwick_2d_rect_go_d61: invariant violated")
)
