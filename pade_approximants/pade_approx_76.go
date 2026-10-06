package pade_approximants

// ComputePadeApprox76 evaluates rational function approximant order 76.
func ComputePadeApprox76(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(1))
}
