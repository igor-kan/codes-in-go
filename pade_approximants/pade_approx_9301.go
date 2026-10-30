package pade_approximants

// ComputePadeApprox9301 evaluates rational function approximant order 9301.
func ComputePadeApprox9301(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
