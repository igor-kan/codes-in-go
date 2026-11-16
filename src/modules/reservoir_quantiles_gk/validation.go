package mod_reservoir_quantiles_gk

import "math"

func ValidateEps(eps float64) error {
	if math.IsNaN(eps) || eps <= 0 || eps >= 0.5 {
		return ErrQuantileInvalid
	}
	return nil
}
