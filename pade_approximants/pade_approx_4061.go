package pade_approximants

// ComputePadeApprox4061 evaluates rational function approximant order 4061.
func ComputePadeApprox4061(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
