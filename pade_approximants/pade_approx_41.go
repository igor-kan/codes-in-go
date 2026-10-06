package pade_approximants

// ComputePadeApprox41 evaluates rational function approximant order 41.
func ComputePadeApprox41(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
