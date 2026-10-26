package pade_approximants

// ComputePadeApprox8981 evaluates rational function approximant order 8981.
func ComputePadeApprox8981(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
