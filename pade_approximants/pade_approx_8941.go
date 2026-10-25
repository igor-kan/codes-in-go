package pade_approximants

// ComputePadeApprox8941 evaluates rational function approximant order 8941.
func ComputePadeApprox8941(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
