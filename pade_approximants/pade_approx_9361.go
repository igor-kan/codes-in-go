package pade_approximants

// ComputePadeApprox9361 evaluates rational function approximant order 9361.
func ComputePadeApprox9361(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
