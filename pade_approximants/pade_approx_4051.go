package pade_approximants

// ComputePadeApprox4051 evaluates rational function approximant order 4051.
func ComputePadeApprox4051(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
