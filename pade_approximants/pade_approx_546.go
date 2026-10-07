package pade_approximants

// ComputePadeApprox546 evaluates rational function approximant order 546.
func ComputePadeApprox546(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(1))
}
