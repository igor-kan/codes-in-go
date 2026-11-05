package pade_approximants

// ComputePadeApprox9731 evaluates rational function approximant order 9731.
func ComputePadeApprox9731(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
