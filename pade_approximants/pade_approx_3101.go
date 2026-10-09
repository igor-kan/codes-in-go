package pade_approximants

// ComputePadeApprox3101 evaluates rational function approximant order 3101.
func ComputePadeApprox3101(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
