package pade_approximants

// ComputePadeApprox7721 evaluates rational function approximant order 7721.
func ComputePadeApprox7721(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
