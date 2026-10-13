package pade_approximants

// ComputePadeApprox7031 evaluates rational function approximant order 7031.
func ComputePadeApprox7031(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
