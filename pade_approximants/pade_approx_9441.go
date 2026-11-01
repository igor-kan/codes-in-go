package pade_approximants

// ComputePadeApprox9441 evaluates rational function approximant order 9441.
func ComputePadeApprox9441(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
