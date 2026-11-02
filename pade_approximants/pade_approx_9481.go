package pade_approximants

// ComputePadeApprox9481 evaluates rational function approximant order 9481.
func ComputePadeApprox9481(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
