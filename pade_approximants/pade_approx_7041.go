package pade_approximants

// ComputePadeApprox7041 evaluates rational function approximant order 7041.
func ComputePadeApprox7041(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
