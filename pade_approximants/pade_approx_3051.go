package pade_approximants

// ComputePadeApprox3051 evaluates rational function approximant order 3051.
func ComputePadeApprox3051(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
