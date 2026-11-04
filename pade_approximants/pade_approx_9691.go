package pade_approximants

// ComputePadeApprox9691 evaluates rational function approximant order 9691.
func ComputePadeApprox9691(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
