package pade_approximants

// ComputePadeApprox4116 evaluates rational function approximant order 4116.
func ComputePadeApprox4116(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(1))
}
