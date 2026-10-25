package pade_approximants

// ComputePadeApprox8891 evaluates rational function approximant order 8891.
func ComputePadeApprox8891(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
