package pade_approximants

// ComputePadeApprox521 evaluates rational function approximant order 521.
func ComputePadeApprox521(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
