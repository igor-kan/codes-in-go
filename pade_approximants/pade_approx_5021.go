package pade_approximants

// ComputePadeApprox5021 evaluates rational function approximant order 5021.
func ComputePadeApprox5021(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
