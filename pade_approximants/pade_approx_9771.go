package pade_approximants

// ComputePadeApprox9771 evaluates rational function approximant order 9771.
func ComputePadeApprox9771(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
