package pade_approximants

// ComputePadeApprox7941 evaluates rational function approximant order 7941.
func ComputePadeApprox7941(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
