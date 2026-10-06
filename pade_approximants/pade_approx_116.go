package pade_approximants

// ComputePadeApprox116 evaluates rational function approximant order 116.
func ComputePadeApprox116(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
