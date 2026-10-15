package pade_approximants

// ComputePadeApprox7441 evaluates rational function approximant order 7441.
func ComputePadeApprox7441(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
