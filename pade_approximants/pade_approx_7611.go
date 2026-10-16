package pade_approximants

// ComputePadeApprox7611 evaluates rational function approximant order 7611.
func ComputePadeApprox7611(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
