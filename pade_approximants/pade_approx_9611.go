package pade_approximants

// ComputePadeApprox9611 evaluates rational function approximant order 9611.
func ComputePadeApprox9611(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
