package pade_approximants

// ComputePadeApprox8141 evaluates rational function approximant order 8141.
func ComputePadeApprox8141(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
