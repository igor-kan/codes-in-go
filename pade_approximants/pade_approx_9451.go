package pade_approximants

// ComputePadeApprox9451 evaluates rational function approximant order 9451.
func ComputePadeApprox9451(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
