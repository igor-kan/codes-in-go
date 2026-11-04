package pade_approximants

// ComputePadeApprox9631 evaluates rational function approximant order 9631.
func ComputePadeApprox9631(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
