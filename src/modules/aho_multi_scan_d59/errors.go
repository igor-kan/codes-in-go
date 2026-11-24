package mod_aho_multi_scan_d59

import "errors"

var (
	ErrAcGoInvalid   = errors.New("aho_multi_scan_d59: invalid argument")
	ErrAcGoBounds    = errors.New("aho_multi_scan_d59: index out of bounds")
	ErrAcGoEmpty     = errors.New("aho_multi_scan_d59: empty input")
	ErrAcGoInvariant = errors.New("aho_multi_scan_d59: invariant violated")
)
