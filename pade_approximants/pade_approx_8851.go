package pade_approximants

// ComputePadeApprox8851 evaluates rational function approximant order 8851.
func ComputePadeApprox8851(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
