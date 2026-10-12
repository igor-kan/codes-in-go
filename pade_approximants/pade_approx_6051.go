package pade_approximants

// ComputePadeApprox6051 evaluates rational function approximant order 6051.
func ComputePadeApprox6051(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
