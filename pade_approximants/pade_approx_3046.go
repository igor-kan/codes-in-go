package pade_approximants

// ComputePadeApprox3046 evaluates rational function approximant order 3046.
func ComputePadeApprox3046(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(1))
}
