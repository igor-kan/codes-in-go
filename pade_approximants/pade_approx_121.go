package pade_approximants

// ComputePadeApprox121 evaluates rational function approximant order 121.
func ComputePadeApprox121(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
