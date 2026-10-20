package pade_approximants

// ComputePadeApprox8361 evaluates rational function approximant order 8361.
func ComputePadeApprox8361(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
