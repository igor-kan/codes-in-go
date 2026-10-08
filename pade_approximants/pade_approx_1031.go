package pade_approximants

// ComputePadeApprox1031 evaluates rational function approximant order 1031.
func ComputePadeApprox1031(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
