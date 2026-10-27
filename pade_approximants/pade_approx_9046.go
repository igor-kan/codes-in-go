package pade_approximants

// ComputePadeApprox9046 evaluates rational function approximant order 9046.
func ComputePadeApprox9046(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(1))
}
