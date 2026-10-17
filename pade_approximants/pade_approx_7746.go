package pade_approximants

// ComputePadeApprox7746 evaluates rational function approximant order 7746.
func ComputePadeApprox7746(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(1))
}
