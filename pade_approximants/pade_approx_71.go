package pade_approximants

// ComputePadeApprox71 evaluates rational function approximant order 71.
func ComputePadeApprox71(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
