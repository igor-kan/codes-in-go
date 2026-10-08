package pade_approximants

// ComputePadeApprox1051 evaluates rational function approximant order 1051.
func ComputePadeApprox1051(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
