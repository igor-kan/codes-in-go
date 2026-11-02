package pade_approximants

// ComputePadeApprox9501 evaluates rational function approximant order 9501.
func ComputePadeApprox9501(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
