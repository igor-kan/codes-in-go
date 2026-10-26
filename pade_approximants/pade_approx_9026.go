package pade_approximants

// ComputePadeApprox9026 evaluates rational function approximant order 9026.
func ComputePadeApprox9026(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
