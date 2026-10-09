package pade_approximants

// ComputePadeApprox3081 evaluates rational function approximant order 3081.
func ComputePadeApprox3081(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
