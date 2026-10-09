package pade_approximants

// ComputePadeApprox3001 evaluates rational function approximant order 3001.
func ComputePadeApprox3001(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
