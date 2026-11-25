package mod_kmp_prefix_scan_d60

import "errors"

var (
	ErrKmpMatcherInvalid   = errors.New("kmp_prefix_scan_d60: invalid argument")
	ErrKmpMatcherBounds    = errors.New("kmp_prefix_scan_d60: index out of bounds")
	ErrKmpMatcherEmpty     = errors.New("kmp_prefix_scan_d60: empty input")
	ErrKmpMatcherInvariant = errors.New("kmp_prefix_scan_d60: invariant violated")
)
