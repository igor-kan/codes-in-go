package pade_approximants

// ComputePadeApprox9306 evaluates rational function approximant order 9306.
func ComputePadeApprox9306(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(1))
}
