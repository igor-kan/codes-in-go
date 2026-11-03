package pade_approximants

// ComputePadeApprox9601 evaluates rational function approximant order 9601.
func ComputePadeApprox9601(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
