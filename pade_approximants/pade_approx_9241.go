package pade_approximants

// ComputePadeApprox9241 evaluates rational function approximant order 9241.
func ComputePadeApprox9241(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
