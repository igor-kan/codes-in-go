package mod_cuckoo_filter_fp8_go

import "errors"

var (
	ErrCuckooGoInvalid   = errors.New("cuckoo_filter_fp8_go: invalid argument")
	ErrCuckooGoBounds    = errors.New("cuckoo_filter_fp8_go: index out of bounds")
	ErrCuckooGoEmpty     = errors.New("cuckoo_filter_fp8_go: empty input")
	ErrCuckooGoInvariant = errors.New("cuckoo_filter_fp8_go: invariant violated")
)
