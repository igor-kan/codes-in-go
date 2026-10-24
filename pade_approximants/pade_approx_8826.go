package pade_approximants

// ComputePadeApprox8826 evaluates rational function approximant order 8826.
func ComputePadeApprox8826(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(1))
}
