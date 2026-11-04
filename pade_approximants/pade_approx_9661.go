package pade_approximants

// ComputePadeApprox9661 evaluates rational function approximant order 9661.
func ComputePadeApprox9661(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
