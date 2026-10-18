package pade_approximants

// ComputePadeApprox7931 evaluates rational function approximant order 7931.
func ComputePadeApprox7931(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
