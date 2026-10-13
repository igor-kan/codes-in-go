package pade_approximants

// ComputePadeApprox7061 evaluates rational function approximant order 7061.
func ComputePadeApprox7061(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
