package pade_approximants

// ComputePadeApprox8876 evaluates rational function approximant order 8876.
func ComputePadeApprox8876(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
