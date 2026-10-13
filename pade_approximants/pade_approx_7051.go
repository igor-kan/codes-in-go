package pade_approximants

// ComputePadeApprox7051 evaluates rational function approximant order 7051.
func ComputePadeApprox7051(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
