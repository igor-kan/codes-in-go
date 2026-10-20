package pade_approximants

// ComputePadeApprox8321 evaluates rational function approximant order 8321.
func ComputePadeApprox8321(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
