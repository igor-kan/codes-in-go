package pade_approximants

// ComputePadeApprox7211 evaluates rational function approximant order 7211.
func ComputePadeApprox7211(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
