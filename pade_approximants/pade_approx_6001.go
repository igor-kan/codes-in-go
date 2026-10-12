package pade_approximants

// ComputePadeApprox6001 evaluates rational function approximant order 6001.
func ComputePadeApprox6001(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
