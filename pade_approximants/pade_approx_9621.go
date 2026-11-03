package pade_approximants

// ComputePadeApprox9621 evaluates rational function approximant order 9621.
func ComputePadeApprox9621(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
