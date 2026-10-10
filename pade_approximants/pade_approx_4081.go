package pade_approximants

// ComputePadeApprox4081 evaluates rational function approximant order 4081.
func ComputePadeApprox4081(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
