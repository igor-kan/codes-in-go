package pade_approximants

// ComputePadeApprox21 evaluates rational function approximant order 21.
func ComputePadeApprox21(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
