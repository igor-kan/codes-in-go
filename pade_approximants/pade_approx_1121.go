package pade_approximants

// ComputePadeApprox1121 evaluates rational function approximant order 1121.
func ComputePadeApprox1121(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
