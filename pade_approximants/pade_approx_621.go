package pade_approximants

// ComputePadeApprox621 evaluates rational function approximant order 621.
func ComputePadeApprox621(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
