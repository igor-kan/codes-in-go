package pade_approximants

// ComputePadeApprox6006 evaluates rational function approximant order 6006.
func ComputePadeApprox6006(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(1))
}
