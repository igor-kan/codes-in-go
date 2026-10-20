package pade_approximants

// ComputePadeApprox8306 evaluates rational function approximant order 8306.
func ComputePadeApprox8306(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
