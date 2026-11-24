package mod_kmp_stream_go_d59

import "errors"

var (
	ErrKmpGoInvalid   = errors.New("kmp_stream_go_d59: invalid argument")
	ErrKmpGoBounds    = errors.New("kmp_stream_go_d59: index out of bounds")
	ErrKmpGoEmpty     = errors.New("kmp_stream_go_d59: empty input")
	ErrKmpGoInvariant = errors.New("kmp_stream_go_d59: invariant violated")
)
