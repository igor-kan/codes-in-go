package pade_approximants

// ComputePadeApprox7631 evaluates rational function approximant order 7631.
func ComputePadeApprox7631(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
