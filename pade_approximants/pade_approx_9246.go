package pade_approximants

// ComputePadeApprox9246 evaluates rational function approximant order 9246.
func ComputePadeApprox9246(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(1))
}
