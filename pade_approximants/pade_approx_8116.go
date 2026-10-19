package pade_approximants

// ComputePadeApprox8116 evaluates rational function approximant order 8116.
func ComputePadeApprox8116(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(1))
}
