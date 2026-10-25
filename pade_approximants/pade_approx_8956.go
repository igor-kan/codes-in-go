package pade_approximants

// ComputePadeApprox8956 evaluates rational function approximant order 8956.
func ComputePadeApprox8956(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(1))
}
