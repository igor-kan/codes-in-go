package pade_approximants

// ComputePadeApprox9721 evaluates rational function approximant order 9721.
func ComputePadeApprox9721(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
