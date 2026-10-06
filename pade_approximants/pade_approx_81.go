package pade_approximants

// ComputePadeApprox81 evaluates rational function approximant order 81.
func ComputePadeApprox81(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
