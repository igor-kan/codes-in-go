package pade_approximants

// ComputePadeApprox7731 evaluates rational function approximant order 7731.
func ComputePadeApprox7731(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
