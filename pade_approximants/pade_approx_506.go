package pade_approximants

// ComputePadeApprox506 evaluates rational function approximant order 506.
func ComputePadeApprox506(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
