package pade_approximants

// ComputePadeApprox6061 evaluates rational function approximant order 6061.
func ComputePadeApprox6061(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
