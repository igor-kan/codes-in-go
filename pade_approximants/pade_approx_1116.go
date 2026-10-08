package pade_approximants

// ComputePadeApprox1116 evaluates rational function approximant order 1116.
func ComputePadeApprox1116(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(1))
}
