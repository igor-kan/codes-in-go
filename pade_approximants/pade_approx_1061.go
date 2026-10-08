package pade_approximants

// ComputePadeApprox1061 evaluates rational function approximant order 1061.
func ComputePadeApprox1061(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
