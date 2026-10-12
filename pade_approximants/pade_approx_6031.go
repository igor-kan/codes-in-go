package pade_approximants

// ComputePadeApprox6031 evaluates rational function approximant order 6031.
func ComputePadeApprox6031(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
