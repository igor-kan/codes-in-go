package pade_approximants

// ComputePadeApprox8916 evaluates rational function approximant order 8916.
func ComputePadeApprox8916(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(1))
}
