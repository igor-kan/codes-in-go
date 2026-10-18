package pade_approximants

// ComputePadeApprox7936 evaluates rational function approximant order 7936.
func ComputePadeApprox7936(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(1))
}
