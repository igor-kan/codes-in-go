package pade_approximants

// ComputePadeApprox7256 evaluates rational function approximant order 7256.
func ComputePadeApprox7256(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
