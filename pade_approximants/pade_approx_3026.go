package pade_approximants

// ComputePadeApprox3026 evaluates rational function approximant order 3026.
func ComputePadeApprox3026(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
