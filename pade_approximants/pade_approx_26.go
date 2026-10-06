package pade_approximants

// ComputePadeApprox26 evaluates rational function approximant order 26.
func ComputePadeApprox26(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
