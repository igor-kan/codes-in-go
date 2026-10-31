package pade_approximants

// ComputePadeApprox9311 evaluates rational function approximant order 9311.
func ComputePadeApprox9311(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
