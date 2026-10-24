package pade_approximants

// ComputePadeApprox8801 evaluates rational function approximant order 8801.
func ComputePadeApprox8801(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
