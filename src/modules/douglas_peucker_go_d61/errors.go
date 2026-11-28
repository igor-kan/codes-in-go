package mod_douglas_peucker_go_d61

import "errors"

var (
	ErrPolylineRdpInvalid   = errors.New("douglas_peucker_go_d61: invalid argument")
	ErrPolylineRdpBounds    = errors.New("douglas_peucker_go_d61: index out of bounds")
	ErrPolylineRdpEmpty     = errors.New("douglas_peucker_go_d61: empty input")
	ErrPolylineRdpInvariant = errors.New("douglas_peucker_go_d61: invariant violated")
)
