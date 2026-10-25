package pade_approximants

// ComputePadeApprox8911 evaluates rational function approximant order 8911.
func ComputePadeApprox8911(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
