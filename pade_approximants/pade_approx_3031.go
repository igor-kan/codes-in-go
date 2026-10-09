package pade_approximants

// ComputePadeApprox3031 evaluates rational function approximant order 3031.
func ComputePadeApprox3031(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
