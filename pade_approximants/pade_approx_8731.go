package pade_approximants

// ComputePadeApprox8731 evaluates rational function approximant order 8731.
func ComputePadeApprox8731(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
