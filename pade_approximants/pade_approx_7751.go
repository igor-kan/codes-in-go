package pade_approximants

// ComputePadeApprox7751 evaluates rational function approximant order 7751.
func ComputePadeApprox7751(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
