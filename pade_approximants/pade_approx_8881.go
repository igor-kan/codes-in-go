package pade_approximants

// ComputePadeApprox8881 evaluates rational function approximant order 8881.
func ComputePadeApprox8881(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
