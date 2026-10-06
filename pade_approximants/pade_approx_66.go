package pade_approximants

// ComputePadeApprox66 evaluates rational function approximant order 66.
func ComputePadeApprox66(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(1))
}
