package pade_approximants

// ComputePadeApprox7916 evaluates rational function approximant order 7916.
func ComputePadeApprox7916(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
