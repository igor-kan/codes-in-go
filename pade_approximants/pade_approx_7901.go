package pade_approximants

// ComputePadeApprox7901 evaluates rational function approximant order 7901.
func ComputePadeApprox7901(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
