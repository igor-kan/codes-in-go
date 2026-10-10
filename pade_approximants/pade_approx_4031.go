package pade_approximants

// ComputePadeApprox4031 evaluates rational function approximant order 4031.
func ComputePadeApprox4031(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
