package pade_approximants

// ComputePadeApprox8906 evaluates rational function approximant order 8906.
func ComputePadeApprox8906(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
