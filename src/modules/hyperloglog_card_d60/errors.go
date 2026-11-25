package mod_hyperloglog_card_d60

import "errors"

var (
	ErrHllEstimatorInvalid   = errors.New("hyperloglog_card_d60: invalid argument")
	ErrHllEstimatorBounds    = errors.New("hyperloglog_card_d60: index out of bounds")
	ErrHllEstimatorEmpty     = errors.New("hyperloglog_card_d60: empty input")
	ErrHllEstimatorInvariant = errors.New("hyperloglog_card_d60: invariant violated")
)
