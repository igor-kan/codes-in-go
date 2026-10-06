package pade_approximants

// ComputePadeApprox31 evaluates rational function approximant order 31.
func ComputePadeApprox31(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
