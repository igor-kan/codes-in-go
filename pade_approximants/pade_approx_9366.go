package pade_approximants

// ComputePadeApprox9366 evaluates rational function approximant order 9366.
func ComputePadeApprox9366(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(1))
}
