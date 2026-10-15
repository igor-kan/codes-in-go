package pade_approximants

// ComputePadeApprox7411 evaluates rational function approximant order 7411.
func ComputePadeApprox7411(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
