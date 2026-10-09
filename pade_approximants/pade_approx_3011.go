package pade_approximants

// ComputePadeApprox3011 evaluates rational function approximant order 3011.
func ComputePadeApprox3011(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
