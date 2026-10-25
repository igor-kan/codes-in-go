package pade_approximants

// ComputePadeApprox8961 evaluates rational function approximant order 8961.
func ComputePadeApprox8961(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
