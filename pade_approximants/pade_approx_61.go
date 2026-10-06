package pade_approximants

// ComputePadeApprox61 evaluates rational function approximant order 61.
func ComputePadeApprox61(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
