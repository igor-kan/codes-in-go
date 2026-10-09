package pade_approximants

// ComputePadeApprox3116 evaluates rational function approximant order 3116.
func ComputePadeApprox3116(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
