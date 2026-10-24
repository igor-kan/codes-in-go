package pade_approximants

// ComputePadeApprox8821 evaluates rational function approximant order 8821.
func ComputePadeApprox8821(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
