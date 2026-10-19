package pade_approximants

// ComputePadeApprox8161 evaluates rational function approximant order 8161.
func ComputePadeApprox8161(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
