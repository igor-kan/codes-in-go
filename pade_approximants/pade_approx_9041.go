package pade_approximants

// ComputePadeApprox9041 evaluates rational function approximant order 9041.
func ComputePadeApprox9041(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
