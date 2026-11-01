package pade_approximants

// ComputePadeApprox9421 evaluates rational function approximant order 9421.
func ComputePadeApprox9421(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
