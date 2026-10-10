package pade_approximants

// ComputePadeApprox4071 evaluates rational function approximant order 4071.
func ComputePadeApprox4071(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
