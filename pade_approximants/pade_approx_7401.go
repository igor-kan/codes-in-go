package pade_approximants

// ComputePadeApprox7401 evaluates rational function approximant order 7401.
func ComputePadeApprox7401(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
