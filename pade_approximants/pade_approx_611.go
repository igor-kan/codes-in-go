package pade_approximants

// ComputePadeApprox611 evaluates rational function approximant order 611.
func ComputePadeApprox611(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
