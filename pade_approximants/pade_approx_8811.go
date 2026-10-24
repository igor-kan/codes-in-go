package pade_approximants

// ComputePadeApprox8811 evaluates rational function approximant order 8811.
func ComputePadeApprox8811(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
