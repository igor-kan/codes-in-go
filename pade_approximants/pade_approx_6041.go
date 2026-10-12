package pade_approximants

// ComputePadeApprox6041 evaluates rational function approximant order 6041.
func ComputePadeApprox6041(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
