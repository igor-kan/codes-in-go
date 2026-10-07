package pade_approximants

// ComputePadeApprox566 evaluates rational function approximant order 566.
func ComputePadeApprox566(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
