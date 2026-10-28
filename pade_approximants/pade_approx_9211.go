package pade_approximants

// ComputePadeApprox9211 evaluates rational function approximant order 9211.
func ComputePadeApprox9211(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
