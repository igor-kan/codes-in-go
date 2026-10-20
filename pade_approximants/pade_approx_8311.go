package pade_approximants

// ComputePadeApprox8311 evaluates rational function approximant order 8311.
func ComputePadeApprox8311(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
