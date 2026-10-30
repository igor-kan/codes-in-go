package pade_approximants

// ComputePadeApprox9281 evaluates rational function approximant order 9281.
func ComputePadeApprox9281(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
