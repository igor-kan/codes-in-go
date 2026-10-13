package pade_approximants

// ComputePadeApprox7046 evaluates rational function approximant order 7046.
func ComputePadeApprox7046(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
