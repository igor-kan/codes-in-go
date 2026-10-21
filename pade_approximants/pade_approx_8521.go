package pade_approximants

// ComputePadeApprox8521 evaluates rational function approximant order 8521.
func ComputePadeApprox8521(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
